package postgres

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go-template/pkg/logger"
)

type DB struct {
	conn   *sql.Conn
	logger *logger.Logger
	DB     *sqlx.DB
}

func New(ctx context.Context, connStr string) *DB {
	db, err := sql.Open("postgres", connStr)
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
		conn: conn,
	}
}
