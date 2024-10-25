package service

import (
	"context"
	"go-template/internal/models/models"
)

type Container interface {
	GetAll(ctx context.Context) ([]models.Container, error)
	Create(ctx context.Context, user models.CreateContainer) (int, error)
	Login(ctx context.Context, user models.CreateContainer) (int, error)
	Upload(ctx context.Context) (int, error)
}
