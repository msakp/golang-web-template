package database

import (
	"context"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/msakp/golang-web-template/internal/config"
	"github.com/msakp/golang-web-template/internal/infrastructure/database/sqlc/storage"
)

type Pg struct {
	queries *storage.Queries
	
	url string
	conn *pgx.Conn
	ctx  context.Context
}

func NewPg(config *config.Config) (*Pg, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, config.PostgresUrl)
	if err != nil {
		return nil, err
	}
	return &Pg{queries: storage.New(conn), ctx: ctx, conn: conn, url: config.PostgresUrl}, nil
}

func (d *Pg) CloseConn() {
	d.conn.Close(d.ctx)

}

func (d *Pg) Migrate() error{
	m, err := migrate.New("file://internal/infrastructure/database/migrations", d.url + "?sslmode=disable")
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange){
		return err
	}
	return nil
	
}

func (d *Pg) Queries() *storage.Queries{
	return d.queries
}

func (d *Pg) Context() context.Context{
	return d.ctx
}
