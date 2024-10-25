package container

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go-template/internal/models/models"
	"go-template/internal/repository"
)

type Container struct {
	db *sqlx.DB
}

func InitContainerRepo(db *sqlx.DB) repository.ContainerRepo {
	return Container{
		db: db,
	}
}

func (u Container) GetAll(ctx context.Context, id int) (*models.Container, error) {
	//TODO implement me
	panic("implement me")
}

func (u Container) Create(ctx context.Context, userCreate models.CreateContainer) (int, error) {
	//TODO implement me
	panic("implement me")
}
