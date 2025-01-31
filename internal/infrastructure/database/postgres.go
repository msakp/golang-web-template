package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/msakp/golang-web-template/internal/infrastructure/config"
	sqlc "github.com/msakp/golang-web-template/internal/infrastructure/database/sqlc/storage"
)

type database struct {
	sqlc.Queries

	db   sqlc.DBTX
	conn *pgx.Conn
	Ctx  context.Context
}

func NewPg(config *config.Config) (*database, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, config.PostgresUrl)
	if err != nil {
		return nil, err
	}
	return &database{db: conn, Ctx: ctx, conn: conn}, nil
}

func (d *database) CloseConn() {
	d.conn.Close(d.Ctx)

}
