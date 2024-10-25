package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"sila-app/pkg/logger"
)

type DB struct {
	logger *logger.Logger
	DB     *sqlx.DB
}

func New(ctx context.Context, connStr string) *DB {
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		panic(err)
	}

	conn, err := db.Conn(ctx)
	if err != nil {
		panic(err)
	}

	_, err = conn.ExecContext(ctx, queryInitContainers)
	if err != nil {
		panic(err)
	}

	return &DB{
		DB: db,
	}
}
