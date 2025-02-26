package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/msakp/golang-web-template/internal/config"
	"github.com/msakp/golang-web-template/internal/database/storage"
	"github.com/msakp/golang-web-template/pkg/logger"
)

const (
	migrationsPath string = "internal/database/migrations"
)

type DB struct {
	queries  *storage.Queries
	connPool *pgxpool.Pool
	url      string
}

func New(ctx context.Context, config *config.Config) *DB {
	pool, err := pgxpool.New(ctx, config.PostgresUrl)
	if err != nil {
		logger.FromCtx(ctx).Error(ctx, fmt.Sprintf("DB Pool error: %s\n", err.Error()))
		return nil
	}

	return &DB{
		queries:  storage.New(),
		connPool: pool,
		url:      config.PostgresUrl,
	}
}

func (d *DB) Migrate(ctx context.Context) {
	m, err := migrate.New(fmt.Sprintf("file://%s", migrationsPath), d.url+"?sslmode=disable")
	if err != nil {
		logger.FromCtx(ctx).Error(ctx, fmt.Sprintf("Migration error: %s", err.Error()))
	}
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.FromCtx(ctx).Error(ctx, fmt.Sprintf("Migration error: %s", err.Error()))
	}
}


func (d *DB) Close(ctx context.Context) {
	d.connPool.Close()
}

func (d *DB) Queries() *storage.Queries {
	return d.queries
}

func (d *DB) Pool() *pgxpool.Pool {
	return d.connPool
}
