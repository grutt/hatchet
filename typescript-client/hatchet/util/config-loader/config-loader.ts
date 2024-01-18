import { parse } from 'yaml';
import { readFileSync } from 'fs';
import * as p from 'path';
import { z } from 'zod';
import { ClientConfig, ClientConfigSchema } from '@clients/client';

type EnvVars =
  | 'HATCHET_CLIENT_TENANT_ID'
  | 'HATCHET_CLIENT_HOST_PORT'
  | 'HATCHET_CLIENT_TLS_CERT_FILE'
  | 'HATCHET_CLIENT_TLS_KEY_FILE'
  | 'HATCHET_CLIENT_TLS_ROOT_CA_FILE'
  | 'HATCHET_CLIENT_TLS_SERVER_NAME';

interface LoadClientConfigOptions {
  path?: string;
}

const DEFAULT_CONFIG_FILE = '.hatchet.yaml';

export class ConfigLoader {
  static load_client_config(config?: LoadClientConfigOptions): Partial<ClientConfig> {
    const yaml = this.load_yaml_config(config?.path);

    return {
      tenant_id: yaml?.tenant_id ?? this.env('HATCHET_CLIENT_TENANT_ID'),
      host_port: yaml?.host_port ?? this.env('HATCHET_CLIENT_HOST_PORT'),
      tls_config: {
        cert_file: yaml?.tls_config?.cert_file ?? this.env('HATCHET_CLIENT_TLS_CERT_FILE')!,
        key_file: yaml?.tls_config?.key_file ?? this.env('HATCHET_CLIENT_TLS_KEY_FILE')!,
        ca_file: yaml?.tls_config?.ca_file ?? this.env('HATCHET_CLIENT_TLS_ROOT_CA_FILE')!,
        server_name: yaml?.tls_config?.server_name ?? this.env('HATCHET_CLIENT_TLS_SERVER_NAME')!,
      },
    };
  }

  static get default_yaml_config_path() {
    return p.join(process.cwd(), DEFAULT_CONFIG_FILE);
  }

  static load_yaml_config(path?: string): ClientConfig | undefined {
    try {
      const configFile = readFileSync(
        p.join(__dirname, path ?? this.default_yaml_config_path),
        'utf8'
      );

      const config = parse(configFile);

      ClientConfigSchema.partial().parse(config);

      return config as ClientConfig;
    } catch (e) {
      if (!path) return undefined;

      if (e instanceof z.ZodError) {
        throw new Error(`Invalid yaml config: ${e.message}`);
      }

      throw e;
    }
  }

  private static env(name: EnvVars): string | undefined {
    return process.env[name];
  }
}
