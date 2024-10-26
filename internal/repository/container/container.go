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
	insertContainer = "INSERT INTO containers(name, link_small, link_big) values ($1, $2, $3) RETURNING id"
	queryUpload     = "UPDATE containers SET document_id = $1 WHERE id = $2"
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
	// Проверка подключения к базе данных
	if err := c.db.PingContext(ctx); err != nil {
		c.logger.Error(ctx, "Failed to connect to database", zap.Error(err))
		return nil, fmt.Errorf("database connection failed: %w", err)
	}

	rows, err := c.db.QueryContext(ctx, queryContainer)
	if err != nil {
		c.logger.Error(ctx, "Failed to execute query", zap.Error(err))
		return nil, err
	}
	for rows.Next() {
		temp := models.Container{}

		if err := rows.Scan(&temp.ID, &temp.Name, &temp.DocumentID, &temp.LinkSmall, &temp.LinkBig); err != nil {
			c.logger.Error(ctx, "failed to scan row", zap.Error(err))
			continue
		}

		out = append(out, temp)
	}

	return out, nil
}

func (c *Container) Create(ctx context.Context, container models.CreateContainer) error {
	c.logger.Info(ctx, fmt.Sprintf("JSON unmarshalled %v", container))
	tx, err := c.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(insertContainer)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(container.Name, container.LinkSmall, container.LinkBig)
	if err != nil {
		return fmt.Errorf("failed to insert data: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (c *Container) Upload(ctx context.Context, documentID int64, containerID int64) (int64, error) {
	tx, err := c.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to start transaction: %w", err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(insertContainer)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(documentID, containerID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert data: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return documentID, nil
}

func (c *Container) UpdateContainer(ctx context.Context, user *models.Container) error {
	//TODO implement me
	panic("implement me")
}
