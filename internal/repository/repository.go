package repository

import (
	"context"
	"go-template/internal/models/models"
)

type ContainerRepo interface {
	Get(ctx context.Context, id int) (*models.Container, error)
	Delete(ctx context.Context, id int) error
	Create(ctx context.Context, userCreate models.CreateContainer) (int, error)
	GetPwdByEmail(ctx context.Context, email string) (string, error)
	GetIDByEmail(ctx context.Context, email string) (int, error)
}
