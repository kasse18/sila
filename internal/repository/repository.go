package repository

import (
	"context"
	"sila-app/internal/models/models"
)

type ContainerRepo interface {
	GetAll(ctx context.Context) ([]models.Container, error)
	Create(ctx context.Context, container models.CreateContainer) error
	UpdateContainer(ctx context.Context, user *models.Container) error
	Upload(ctx context.Context, documentID int64, containerID int64) (int64, error)
}
