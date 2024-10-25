package container

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go-template/internal/models/models"
	"go-template/pkg/logger"
	"go.uber.org/zap"
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

func InitContainerRepo(db *sqlx.DB) *Container {
	return &Container{
		db: db,
	}
}

func (c *Container) GetAll(ctx context.Context) ([]models.Container, error) {
	out := []models.Container{}
	rows, err := c.db.QueryContext(ctx, queryContainer)
	if err != nil {
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

func (c *Container) Create(ctx context.Context, containerCreate models.CreateContainer) (int, error) {
	tx, err := c.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}
	var id int

	newContainer := models.CreateContainer{
		containerCreate.Name,
		containerCreate.LinkSmall,
		containerCreate.LinkBig,
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
