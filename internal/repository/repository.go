package repository

import (
	"context"
	"go-template/internal/models/models"
)

type ContainerRepo interface {
	GetAll(ctx context.Context) ([]models.Container, error)
	Create(ctx context.Context, userCreate models.CreateContainer) (int, error)
	UpdateContainer(ctx context.Context, user *models.Container) error
}
