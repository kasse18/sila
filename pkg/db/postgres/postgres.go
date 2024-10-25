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
	GetUser(ctx context.Context, user *models.User) error
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, user *models.User) error
}

func (r *DB) GetContainer(ctx context.Context, user *models.User) error {
	row := r.DB.QueryRowContext(ctx, queryContainer, user.ID)

	if err := row.Scan(&user.ID, &user.Username); err != nil {
		return err
	}

	return nil
}

func (r *DB) CreateContainer(ctx context.Context, user *models.User) error {
	_, err := r.DB.ExecContext(ctx, insertContainer, user.ID, user.Username)
	if err != nil {
		return err
	}

	return nil
}
