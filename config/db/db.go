package db

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type ConfigDB struct {
	URI string
}

func NewConfigDB(uri string) *ConfigDB {
	return &ConfigDB{
		URI: uri,
	}
}

// OpenDB Inicia una conexión con la uri principal.
func (c *ConfigDB) OpenDB(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, c.URI)
}
