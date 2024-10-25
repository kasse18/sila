package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go-template/internal/models/models"
	"go-template/pkg/logger"
)

type DB struct {
	DB     *sqlx.DB
	logger *logger.Logger
}

func New(connStr string) *DB {
	dsn := fmt.Sprintf(connStr)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err)
	}
	if _, err := db.Conn(context.Background()); err != nil {
		panic(err)
	}
	return &DB{DB: db}
}

type Client interface {
	GetContainer(ctx context.Context, user *models.Container) error
	CreateContainer(ctx context.Context, user *models.Container) error
	UpdateContainer(ctx context.Context, user *models.Container) error
}

func (r *DB) GetContainer(ctx context.Context, container *models.Container) error {
	row := r.DB.QueryRowContext(ctx, queryContainer, container.ID)

	if err := row.Scan(&container.ID, &container.Name); err != nil {
		return err
	}

	return nil
}

func (r *DB) CreateContainer(ctx context.Context, container *models.Container) error {
	_, err := r.DB.ExecContext(ctx, insertContainer, container.ID, container.Name, container.LinkSmall, container.LinkBig)
	if err != nil {
		return err
	}

	return nil
}
