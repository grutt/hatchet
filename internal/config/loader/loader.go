// Adapted from: https://github.com/hatchet-dev/hatchet-v1-archived/blob/3c2c13168afa1af68d4baaf5ed02c9d49c5f0323/internal/config/loader/loader.go

package loader

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/hatchet-dev/hatchet/internal/auth/cookie"
	"github.com/hatchet-dev/hatchet/internal/auth/oauth"
	"github.com/hatchet-dev/hatchet/internal/auth/token"
	"github.com/hatchet-dev/hatchet/internal/config/database"
	"github.com/hatchet-dev/hatchet/internal/config/loader/loaderutils"
	"github.com/hatchet-dev/hatchet/internal/config/server"
	"github.com/hatchet-dev/hatchet/internal/encryption"
	"github.com/hatchet-dev/hatchet/internal/logger"
	"github.com/hatchet-dev/hatchet/internal/repository/prisma"
	"github.com/hatchet-dev/hatchet/internal/repository/prisma/db"
	"github.com/hatchet-dev/hatchet/internal/services/ingestor"
	"github.com/hatchet-dev/hatchet/internal/taskqueue/rabbitmq"
	"github.com/hatchet-dev/hatchet/internal/validator"
)

// LoadDatabaseConfigFile loads the database config file via viper
func LoadDatabaseConfigFile(files ...[]byte) (*database.ConfigFile, error) {
	configFile := &database.ConfigFile{}
	f := database.BindAllEnv

	_, err := loaderutils.LoadConfigFromViper(f, configFile, files...)

	return configFile, err
}

// LoadServerConfigFile loads the server config file via viper
func LoadServerConfigFile(files ...[]byte) (*server.ServerConfigFile, error) {
	configFile := &server.ServerConfigFile{}
	f := server.BindAllEnv

	_, err := loaderutils.LoadConfigFromViper(f, configFile, files...)

	return configFile, err
}

type ConfigLoader struct {
	directory string
}

func NewConfigLoader(directory string) *ConfigLoader {
	return &ConfigLoader{directory}
}

// LoadDatabaseConfig loads the database configuration
func (c *ConfigLoader) LoadDatabaseConfig() (res *database.Config, err error) {
	sharedFilePath := filepath.Join(c.directory, "database.yaml")
	configFileBytes, err := loaderutils.GetConfigBytes(sharedFilePath)

	if err != nil {
		return nil, err
	}

	cf, err := LoadDatabaseConfigFile(configFileBytes...)

	if err != nil {
		return nil, err
	}

	return GetDatabaseConfigFromConfigFile(cf)
}

// LoadServerConfig loads the server configuration
func (c *ConfigLoader) LoadServerConfig() (res *server.ServerConfig, err error) {
	sharedFilePath := filepath.Join(c.directory, "server.yaml")
	configFileBytes, err := loaderutils.GetConfigBytes(sharedFilePath)

	if err != nil {
		return nil, err
	}

	dc, err := c.LoadDatabaseConfig()

	if err != nil {
		return nil, err
	}

	cf, err := LoadServerConfigFile(configFileBytes...)

	if err != nil {
		return nil, err
	}

	return GetServerConfigFromConfigfile(dc, cf)
}

func GetDatabaseConfigFromConfigFile(cf *database.ConfigFile) (res *database.Config, err error) {
	l := logger.NewStdErr(&cf.Logger, "database")

	databaseUrl := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		cf.PostgresUsername,
		cf.PostgresPassword,
		cf.PostgresHost,
		cf.PostgresPort,
		cf.PostgresDbName,
		cf.PostgresSSLMode,
	)

	os.Setenv("DATABASE_URL", databaseUrl)

	client := db.NewClient(
	// db.WithDatasourceURL(databaseUrl),
	)

	if err := client.Prisma.Connect(); err != nil {
		return nil, err
	}

	config, err := pgxpool.ParseConfig(databaseUrl)

	if err != nil {
		return nil, err
	}

	config.MaxConns = 20

	pool, err := pgxpool.NewWithConfig(context.Background(), config)

	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}

	return &database.Config{
		Disconnect: client.Prisma.Disconnect,
		Repository: prisma.NewPrismaRepository(client, pool, prisma.WithLogger(&l)),
		Seed:       cf.Seed,
	}, nil
}

func GetServerConfigFromConfigfile(dc *database.Config, cf *server.ServerConfigFile) (res *server.ServerConfig, err error) {
	l := logger.NewStdErr(&cf.Logger, "server")

	tls, err := loaderutils.LoadServerTLSConfig(&cf.TLS)

	if err != nil {
		return nil, fmt.Errorf("could not load TLS config: %w", err)
	}

	ss, err := cookie.NewUserSessionStore(
		cookie.WithSessionRepository(dc.Repository.UserSession()),
		cookie.WithCookieAllowInsecure(cf.Auth.Cookie.Insecure),
		cookie.WithCookieDomain(cf.Auth.Cookie.Domain),
		cookie.WithCookieName(cf.Auth.Cookie.Name),
		cookie.WithCookieSecrets(getStrArr(cf.Auth.Cookie.Secrets)...),
	)

	if err != nil {
		return nil, fmt.Errorf("could not create session store: %w", err)
	}

	tq := rabbitmq.New(
		context.Background(),
		rabbitmq.WithURL(cf.TaskQueue.RabbitMQ.URL),
		rabbitmq.WithLogger(&l),
	)

	ingestor, err := ingestor.NewIngestor(
		ingestor.WithEventRepository(dc.Repository.Event()),
		ingestor.WithTaskQueue(tq),
	)

	if err != nil {
		return nil, fmt.Errorf("could not create ingestor: %w", err)
	}

	auth := server.AuthConfig{
		ConfigFile: cf.Auth,
	}

	if cf.Auth.Google.Enabled {
		if cf.Auth.Google.ClientID == "" {
			return nil, fmt.Errorf("google client id is required")
		}

		if cf.Auth.Google.ClientSecret == "" {
			return nil, fmt.Errorf("google client secret is required")
		}

		gClient := oauth.NewGoogleClient(&oauth.Config{
			ClientID:     cf.Auth.Google.ClientID,
			ClientSecret: cf.Auth.Google.ClientSecret,
			BaseURL:      cf.Runtime.ServerURL,
			Scopes:       cf.Auth.Google.Scopes,
		})

		auth.GoogleOAuthConfig = gClient
	}

	encryptionSvc, err := loadEncryptionSvc(cf)

	if err != nil {
		return nil, fmt.Errorf("could not load encryption service: %w", err)
	}

	// create a new JWT manager
	auth.JWTManager, err = token.NewJWTManager(encryptionSvc, dc.Repository.APIToken(), &token.TokenOpts{
		Issuer:               cf.Runtime.ServerURL,
		Audience:             cf.Runtime.ServerURL,
		GRPCBroadcastAddress: cf.Runtime.GRPCBroadcastAddress,
		ServerURL:            cf.Runtime.ServerURL,
	})

	if err != nil {
		return nil, fmt.Errorf("could not create JWT manager: %w", err)
	}

	return &server.ServerConfig{
		Runtime:       cf.Runtime,
		Auth:          auth,
		Encryption:    encryptionSvc,
		Config:        dc,
		TaskQueue:     tq,
		Services:      cf.Services,
		Logger:        &l,
		TLSConfig:     tls,
		SessionStore:  ss,
		Validator:     validator.NewDefaultValidator(),
		Ingestor:      ingestor,
		OpenTelemetry: cf.OpenTelemetry,
	}, nil
}

func getStrArr(v string) []string {
	return strings.Split(v, " ")
}

func loadEncryptionSvc(cf *server.ServerConfigFile) (encryption.EncryptionService, error) {
	var err error

	hasLocalMasterKeyset := cf.Encryption.MasterKeyset != "" || cf.Encryption.MasterKeysetFile != ""
	isCloudKMSEnabled := cf.Encryption.CloudKMS.Enabled

	if !hasLocalMasterKeyset && !isCloudKMSEnabled {
		return nil, fmt.Errorf("encryption is required")
	}

	if hasLocalMasterKeyset && isCloudKMSEnabled {
		return nil, fmt.Errorf("cannot use both encryption and cloud kms")
	}

	hasJWTKeys := (cf.Encryption.JWT.PublicJWTKeyset != "" || cf.Encryption.JWT.PublicJWTKeysetFile != "") &&
		(cf.Encryption.JWT.PrivateJWTKeyset != "" || cf.Encryption.JWT.PrivateJWTKeysetFile != "")

	if !hasJWTKeys {
		return nil, fmt.Errorf("jwt encryption is required")
	}

	privateJWT := cf.Encryption.JWT.PrivateJWTKeyset

	if cf.Encryption.JWT.PrivateJWTKeysetFile != "" {
		privateJWTBytes, err := loaderutils.GetFileBytes(cf.Encryption.JWT.PrivateJWTKeysetFile)

		if err != nil {
			return nil, fmt.Errorf("could not load private jwt keyset file: %w", err)
		}

		privateJWT = string(privateJWTBytes)
	}

	publicJWT := cf.Encryption.JWT.PublicJWTKeyset

	if cf.Encryption.JWT.PublicJWTKeysetFile != "" {
		publicJWTBytes, err := loaderutils.GetFileBytes(cf.Encryption.JWT.PublicJWTKeysetFile)

		if err != nil {
			return nil, fmt.Errorf("could not load public jwt keyset file: %w", err)
		}

		publicJWT = string(publicJWTBytes)
	}

	var encryptionSvc encryption.EncryptionService

	if hasLocalMasterKeyset {
		masterKeyset := cf.Encryption.MasterKeyset

		if cf.Encryption.MasterKeysetFile != "" {
			masterKeysetBytes, err := loaderutils.GetFileBytes(cf.Encryption.MasterKeysetFile)

			if err != nil {
				return nil, fmt.Errorf("could not load master keyset file: %w", err)
			}

			masterKeyset = string(masterKeysetBytes)
		}

		encryptionSvc, err = encryption.NewLocalEncryption(
			[]byte(masterKeyset),
			[]byte(privateJWT),
			[]byte(publicJWT),
		)

		if err != nil {
			return nil, fmt.Errorf("could not create raw keyset encryption service: %w", err)
		}
	}

	if isCloudKMSEnabled {
		encryptionSvc, err = encryption.NewCloudKMSEncryption(
			cf.Encryption.CloudKMS.KeyURI,
			[]byte(cf.Encryption.CloudKMS.CredentialsJSON),
			[]byte(privateJWT),
			[]byte(publicJWT),
		)

		if err != nil {
			return nil, fmt.Errorf("could not create CloudKMS encryption service: %w", err)
		}
	}

	return encryptionSvc, nil
}
