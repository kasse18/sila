package container

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"sila-app/internal/models/models"
	"sila-app/pkg/logger"
)

const (
	queryContainer  = "SELECT * FROM containers"
	insertContainer = "INSERT INTO containers(id, name, link_small, link_big) values ($1, $2, $3, $4) RETURNING id"
	queryUpload     = "UPDATE containers SET document_id = $1"
)

type Container struct {
	db     *sqlx.DB
	logger logger.Logger
}

func InitContainerRepo(db *sqlx.DB, logger *logger.Logger) *Container {
	return &Container{
		db:     db,
		logger: *logger,
	}
}

func (c *Container) GetAll(ctx context.Context) ([]models.Container, error) {
	c.logger.Info(ctx, "Starting GetAll operation")
	out := []models.Container{}

	err := c.db.PingContext(ctx)
	if err != nil {
		c.logger.Error(ctx, "Failed to connect to database", zap.Error(err))
		return nil, fmt.Errorf("database connection failed: %w", err)
	}

	rows, err := c.db.QueryxContext(ctx, queryContainer)
	if err != nil {
		c.logger.Error(ctx, "Failed to execute query", zap.Error(err))
		return nil, err
	}

	for rows.Next() {
		temp := models.Container{}

		if err := rows.StructScan(&temp); err != nil {
			c.logger.Error(ctx, "failed to scan row", zap.Error(err))
			continue
		}

		out = append(out, temp)
	}

	return out, nil
}

func (c *Container) Create(ctx context.Context, containerCreate models.CreateContainer) (int, error) {
	tx, err := c.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}
	var id int

	newContainer := models.CreateContainer{
		Name:      containerCreate.Name,
		LinkSmall: containerCreate.LinkSmall,
		LinkBig:   containerCreate.LinkBig,
	}
	row := c.db.QueryRowContext(ctx, insertContainer, newContainer.Name, newContainer.LinkSmall, newContainer.LinkBig)
	err = row.Scan(&id)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return 0, rbErr
		}
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return id, nil
}

func (c *Container) UpdateContainer(ctx context.Context, container *models.Container) error {
	//TODO implement me

	panic("implement me")
}
