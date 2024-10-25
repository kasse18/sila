package repository

import (
	"context"
	"go-template/internal/models/models"
)

type ContainerRepo interface {
	GetAll(ctx context.Context, id int) (*models.Container, error)
	Create(ctx context.Context, userCreate models.CreateContainer) (int, error)
}
