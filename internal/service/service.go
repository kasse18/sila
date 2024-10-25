package service

import (
	"context"
	"go-template/internal/models/models"
)

type User interface {
	GetMe(ctx context.Context, id int) (*models.User, error)
	Delete(ctx context.Context, id int) error
	Create(ctx context.Context, user models.CreateUser) (int, error)
	Login(ctx context.Context, user models.CreateUser) (int, error)
}
