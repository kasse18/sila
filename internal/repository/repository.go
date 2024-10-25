package repository

import (
	"context"
	"sila-app/internal/models/models"
)

type ContainerRepo interface {
	GetAll(ctx context.Context) ([]models.Container, error)
	Create(ctx context.Context, jsonData []byte) error
	UpdateContainer(ctx context.Context, user *models.Container) error
}
