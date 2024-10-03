package prisma

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"

	"github.com/hatchet-dev/hatchet/internal/telemetry"
	"github.com/hatchet-dev/hatchet/pkg/repository"
	"github.com/hatchet-dev/hatchet/pkg/repository/metered"
	"github.com/hatchet-dev/hatchet/pkg/repository/prisma/db"
	"github.com/hatchet-dev/hatchet/pkg/repository/prisma/dbsqlc"
	"github.com/hatchet-dev/hatchet/pkg/repository/prisma/sqlchelpers"
	"github.com/hatchet-dev/hatchet/pkg/validator"
)

type eventAPIRepository struct {
	client  *db.PrismaClient
	pool    *pgxpool.Pool
	v       validator.Validator
	queries *dbsqlc.Queries
	l       *zerolog.Logger
}

func NewEventAPIRepository(client *db.PrismaClient, pool *pgxpool.Pool, v validator.Validator, l *zerolog.Logger) repository.EventAPIRepository {
	queries := dbsqlc.New()

	return &eventAPIRepository{
		client:  client,
		pool:    pool,
		v:       v,
		queries: queries,
		l:       l,
	}
}

func (r *eventAPIRepository) ListEvents(ctx context.Context, tenantId string, opts *repository.ListEventOpts) (*repository.ListEventResult, error) {
	if err := r.v.Validate(opts); err != nil {
		return nil, err
	}

	res := &repository.ListEventResult{}

	pgTenantId := &pgtype.UUID{}

	if err := pgTenantId.Scan(tenantId); err != nil {
		return nil, err
	}

	queryParams := dbsqlc.ListEventsParams{
		TenantId: *pgTenantId,
	}

	countParams := dbsqlc.CountEventsParams{
		TenantId: *pgTenantId,
	}

	if opts.Ids != nil {
		queryParams.EventIds = make([]pgtype.UUID, len(opts.Ids))
		countParams.EventIds = make([]pgtype.UUID, len(opts.Ids))

		for i := range opts.Ids {
			queryParams.EventIds[i] = sqlchelpers.UUIDFromStr(opts.Ids[i])
			countParams.EventIds[i] = sqlchelpers.UUIDFromStr(opts.Ids[i])
		}
	}

	if opts.Search != nil {
		queryParams.Search = sqlchelpers.TextFromStr(*opts.Search)
		countParams.Search = sqlchelpers.TextFromStr(*opts.Search)
	}

	if opts.Offset != nil {
		queryParams.Offset = *opts.Offset
	}

	if opts.Limit != nil {
		queryParams.Limit = *opts.Limit
	}

	if opts.Keys != nil {
		queryParams.Keys = opts.Keys
		countParams.Keys = opts.Keys
	}

	if opts.Workflows != nil {
		queryParams.Workflows = opts.Workflows
		countParams.Workflows = opts.Workflows
	}

	if opts.WorkflowRunStatus != nil {
		statuses := make([]string, 0)

		for _, status := range opts.WorkflowRunStatus {
			statuses = append(statuses, string(status))
		}

		queryParams.Statuses = statuses
		countParams.Statuses = statuses
	}

	if opts.AdditionalMetadata != nil {
		queryParams.AdditionalMetadata = opts.AdditionalMetadata
		countParams.AdditionalMetadata = opts.AdditionalMetadata
	}

	orderByField := "createdAt"
	orderByDirection := "DESC"

	if opts.OrderBy != nil {
		orderByField = *opts.OrderBy
	}

	if opts.OrderDirection != nil {
		orderByDirection = *opts.OrderDirection
	}

	queryParams.Orderby = orderByField + " " + orderByDirection
	countParams.Orderby = orderByField + " " + orderByDirection

	tx, err := r.pool.Begin(context.Background())

	if err != nil {
		return nil, err
	}

	defer deferRollback(context.Background(), r.l, tx.Rollback)

	events, err := r.queries.ListEvents(ctx, tx, queryParams)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			events = make([]*dbsqlc.ListEventsRow, 0)
		} else {
			return nil, fmt.Errorf("could not list events: %w", err)
		}
	}

	count, err := r.queries.CountEvents(ctx, tx, countParams)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			count = 0
		} else {
			return nil, fmt.Errorf("could not count events: %w", err)
		}
	}

	err = tx.Commit(ctx)

	if err != nil {
		return nil, fmt.Errorf("could not commit transaction: %w", err)
	}

	res.Rows = events
	res.Count = int(count)

	return res, nil
}

func (r *eventAPIRepository) ListEventKeys(tenantId string) ([]string, error) {
	var rows []struct {
		Key string `json:"key"`
	}

	err := r.client.Prisma.QueryRaw(
		`
		SELECT DISTINCT ON("Event"."key") "Event"."key"
		FROM "Event"
		WHERE
		"Event"."tenantId"::text = $1
		`,
		tenantId,
	).Exec(context.Background(), &rows)

	if err != nil {
		return nil, err
	}

	keys := make([]string, len(rows))

	for i, row := range rows {
		keys[i] = row.Key
	}

	return keys, nil
}

func sizeOfEvent(item *repository.CreateEventOpts) int {
	return len(item.Data) + len(item.AdditionalMetadata)
}

func (e *eventEngineRepository) startBufferLoop() error {

	tenantBufOpts := TenantBufManagerOpts[*repository.CreateEventOpts, *dbsqlc.Event]{OutputFunc: e.BulkCreateEventSharedTenant, SizeFunc: sizeOfEvent, L: e.l, V: e.v}
	var err error
	e.bulkCreateBuffer, err = NewTenantBufManager(tenantBufOpts)

	return err

}

func (r *eventAPIRepository) GetEventById(id string) (*db.EventModel, error) {
	return r.client.Event.FindFirst(
		db.Event.ID.Equals(id),
		db.Event.DeletedAt.IsNull(),
	).Exec(context.Background())
}

func (r *eventAPIRepository) ListEventsById(tenantId string, ids []string) ([]db.EventModel, error) {
	return r.client.Event.FindMany(
		db.Event.ID.In(ids),
		db.Event.TenantID.Equals(tenantId),
		db.Event.DeletedAt.IsNull(),
	).Exec(context.Background())
}

type eventEngineRepository struct {
	pool             *pgxpool.Pool
	v                validator.Validator
	queries          *dbsqlc.Queries
	l                *zerolog.Logger
	m                *metered.Metered
	bulkCreateBuffer *TenantBufferManager[*repository.CreateEventOpts, *dbsqlc.Event]
	callbacks        []repository.Callback[*dbsqlc.Event]
}

func (r *eventEngineRepository) cleanup() error {
	return r.bulkCreateBuffer.cleanup()
}

func NewEventEngineRepository(pool *pgxpool.Pool, v validator.Validator, l *zerolog.Logger, m *metered.Metered) (repository.EventEngineRepository, func() error, error) {
	queries := dbsqlc.New()

	e := eventEngineRepository{
		pool:    pool,
		v:       v,
		queries: queries,
		l:       l,
		m:       m,
	}
	err := e.startBufferLoop()

	return &e, e.cleanup, err
}

func (r *eventEngineRepository) RegisterCreateCallback(callback repository.Callback[*dbsqlc.Event]) {
	if r.callbacks == nil {
		r.callbacks = make([]repository.Callback[*dbsqlc.Event], 0)
	}

	r.callbacks = append(r.callbacks, callback)
}

func (r *eventEngineRepository) GetEventForEngine(ctx context.Context, tenantId, id string) (*dbsqlc.Event, error) {
	return r.queries.GetEventForEngine(ctx, r.pool, sqlchelpers.UUIDFromStr(id))
}

func (r *eventEngineRepository) CreateEvent(ctx context.Context, opts *repository.CreateEventOpts) (*dbsqlc.Event, error) {
	return metered.MakeMetered(ctx, r.m, dbsqlc.LimitResourceEVENT, opts.TenantId, 1, func() (*string, *dbsqlc.Event, error) {

		_, span := telemetry.NewSpan(ctx, "db-create-event")
		defer span.End()

		if err := r.v.Validate(opts); err != nil {
			return nil, nil, err
		}

		createOpts := repository.CreateEventOpts{
			TenantId:           opts.TenantId,
			Key:                opts.Key,
			Data:               opts.Data,
			AdditionalMetadata: opts.AdditionalMetadata,
			ReplayedEvent:      opts.ReplayedEvent,
		}

		done, err := r.bulkCreateBuffer.BuffItem(opts.TenantId, &createOpts)

		if err != nil {
			return nil, nil, fmt.Errorf("could not buffer event: %w", err)
		}
		var response *flushResponse[*dbsqlc.Event]

		select {
		case response = <-done:
		case <-ctx.Done():
			return nil, nil, ctx.Err()
		case <-time.After(20 * time.Second):
			return nil, nil, fmt.Errorf("timeout waiting for event to be flushed to db")
		}

		if response.err != nil {
			return nil, nil, fmt.Errorf("could not create event: %w", response.err)
		}

		e := response.result

		for _, cb := range r.callbacks {
			cb.Do(r.l, opts.TenantId, e)
		}

		id := sqlchelpers.UUIDToStr(e.ID)

		if e.TenantId != sqlchelpers.UUIDFromStr(opts.TenantId) {
			panic("tenant id mismatch")
		}

		return &id, e, nil
	})
}
func (r *eventEngineRepository) BulkCreateEvent(ctx context.Context, opts *repository.BulkCreateEventOpts) (*repository.BulkCreateEventResult, error) {

	numberOfResources := len(opts.Events)
	if numberOfResources < math.MinInt32 || numberOfResources > math.MaxInt32 {
		return nil, fmt.Errorf("number of resources is out of range")
	}

	return metered.MakeMetered(ctx, r.m, dbsqlc.LimitResourceEVENT, opts.TenantId, int32(numberOfResources), func() (*string, *repository.BulkCreateEventResult, error) {

		ctx, span := telemetry.NewSpan(ctx, "db-bulk-create-event")
		defer span.End()

		if err := r.v.Validate(opts); err != nil {

			return nil, nil, err
		}
		params := make([]dbsqlc.CreateEventsParams, len(opts.Events))

		for i, event := range opts.Events {

			params[i] = dbsqlc.CreateEventsParams{
				ID:                 sqlchelpers.UUIDFromStr(uuid.New().String()),
				Key:                event.Key,
				TenantId:           sqlchelpers.UUIDFromStr(event.TenantId),
				Data:               event.Data,
				AdditionalMetadata: event.AdditionalMetadata,
			}

			if event.ReplayedEvent != nil {
				params[i].ReplayedFromId = sqlchelpers.UUIDFromStr(*event.ReplayedEvent)
			}

		}

		// start a transaction
		tx, err := r.pool.Begin(ctx)

		if err != nil {
			return nil, nil, err
		}

		defer deferRollback(ctx, r.l, tx.Rollback)

		insertCount, err := r.queries.CreateEvents(
			ctx,
			tx,
			params,
		)

		if err != nil {
			return nil, nil, fmt.Errorf("could not create events: %w", err)
		}

		r.l.Info().Msgf("inserted %d events", insertCount)

		events, err := r.queries.GetInsertedEvents(ctx, tx)

		if err != nil {
			return nil, nil, fmt.Errorf("could not retrieve inserted events: %w", err)
		}
		err = tx.Commit(ctx)

		if err != nil {
			return nil, nil, fmt.Errorf("could not commit transaction: %w", err)
		}

		var returnString string

		for _, e := range events {

			for _, cb := range r.callbacks {
				cb.Do(r.l, opts.TenantId, e)
			}

		}

		if len(events) > 0 {

			returnString = sqlchelpers.UUIDToStr(events[0].ID)
		}

		// TODO is this return string important?
		return &returnString, &repository.BulkCreateEventResult{Events: events}, nil
	})
}
func (r *eventEngineRepository) BulkCreateEventSharedTenant(ctx context.Context, opts []*repository.CreateEventOpts) ([]*dbsqlc.Event, error) {

	// need to do the metering beforehand
	numberOfResources := len(opts)
	if numberOfResources < math.MinInt32 || numberOfResources > math.MaxInt32 {
		return nil, fmt.Errorf("number of resources is out of range")
	}

	ctx, span := telemetry.NewSpan(ctx, "db-bulk-create-event-shared-tenant")
	defer span.End()

	for _, opt := range opts {

		if err := r.v.Validate(opt); err != nil {
			return nil, err
		}
	}
	params := make([]dbsqlc.CreateEventsParams, len(opts))

	for i, event := range opts {

		if i > math.MaxInt32 || i < math.MinInt32 {
			return nil, fmt.Errorf("number of resources is out of range for int 32")
		}

		params[i] = dbsqlc.CreateEventsParams{
			ID:                 sqlchelpers.UUIDFromStr(uuid.New().String()),
			Key:                event.Key,
			TenantId:           sqlchelpers.UUIDFromStr(event.TenantId),
			Data:               event.Data,
			AdditionalMetadata: event.AdditionalMetadata,
			InsertOrder:        sqlchelpers.ToInt(int32(i)),
		}

		if event.ReplayedEvent != nil {
			params[i].ReplayedFromId = sqlchelpers.UUIDFromStr(*event.ReplayedEvent)
		}

	}

	// start a transaction
	tx, err := r.pool.Begin(ctx)

	if err != nil {
		return nil, err
	}

	defer deferRollback(ctx, r.l, tx.Rollback)

	insertCount, err := r.queries.CreateEvents(
		ctx,
		tx,
		params,
	)

	if err != nil {
		return nil, fmt.Errorf("could not create events: %w", err)
	}

	r.l.Info().Msgf("inserted %d events", insertCount)

	events, err := r.queries.GetInsertedEvents(ctx, tx)

	if err != nil {
		return nil, fmt.Errorf("could not retrieve inserted events: %w", err)
	}
	err = tx.Commit(ctx)

	if err != nil {
		return nil, fmt.Errorf("could not commit transaction: %w", err)
	}

	for _, e := range events {

		tenantId := sqlchelpers.UUIDToStr(e.TenantId)

		for _, cb := range r.callbacks {
			cb.Do(r.l, tenantId, e)
		}

	}

	// TODO is this return string important?
	return events, nil

}

func (r *eventEngineRepository) ListEventsByIds(ctx context.Context, tenantId string, ids []string) ([]*dbsqlc.Event, error) {
	pgIds := make([]pgtype.UUID, len(ids))

	for i, id := range ids {
		if err := pgIds[i].Scan(id); err != nil {
			return nil, err
		}
	}

	pgTenantId := sqlchelpers.UUIDFromStr(tenantId)

	return r.queries.ListEventsByIDs(ctx, r.pool, dbsqlc.ListEventsByIDsParams{
		Tenantid: pgTenantId,
		Ids:      pgIds,
	})
}

func (r *eventEngineRepository) SoftDeleteExpiredEvents(ctx context.Context, tenantId string, before time.Time) (bool, error) {
	hasMore, err := r.queries.SoftDeleteExpiredEvents(ctx, r.pool, dbsqlc.SoftDeleteExpiredEventsParams{
		Tenantid:      sqlchelpers.UUIDFromStr(tenantId),
		Createdbefore: sqlchelpers.TimestampFromTime(before),
		Limit:         1000,
	})

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	return hasMore, nil
}

func (r *eventEngineRepository) ClearEventPayloadData(ctx context.Context, tenantId string) (bool, error) {
	hasMore, err := r.queries.ClearEventPayloadData(ctx, r.pool, dbsqlc.ClearEventPayloadDataParams{
		Tenantid: sqlchelpers.UUIDFromStr(tenantId),
		Limit:    1000,
	})

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	return hasMore, nil
}
