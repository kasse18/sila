package user

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go-template/internal/models/models"
	"go-template/internal/repository"
)

type Container struct {
	db *sqlx.DB
}

func (u Container) Get(ctx context.Context, id int) (*models.Container, error) {
	//TODO implement me
	panic("implement me")
}

func (u Container) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (u Container) Create(ctx context.Context, userCreate models.CreateContainer) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (u Container) GetPwdByEmail(ctx context.Context, email string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (u Container) GetIDByEmail(ctx context.Context, email string) (int, error) {
	//TODO implement me
	panic("implement me")
}

func InitUserRepo(db *sqlx.DB) repository.UserRepo {
	return Container{
		db: db,
	}
}
