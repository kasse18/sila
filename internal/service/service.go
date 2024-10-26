package service

import (
	"context"
	"sila-app/internal/models/models"
)

type Container interface {
	GetAll(ctx context.Context) ([]models.Container, error)
	Create(ctx context.Context, container models.CreateContainer) error
	Login(ctx context.Context, user models.CreateContainer) (int, error)
	Upload(ctx context.Context, documentID int64, containerID int64) (int, error)
}
