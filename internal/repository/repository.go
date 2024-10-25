package repository

import (
	"context"
	"go-template/internal/models/models"
)

type UserRepo interface {
	Get(ctx context.Context, id int) (*models.User, error)
	Delete(ctx context.Context, id int) error
	Create(ctx context.Context, userCreate models.CreateUser) (int, error)
	GetPwdByEmail(ctx context.Context, email string) (string, error)
	GetIDByEmail(ctx context.Context, email string) (int, error)
}
