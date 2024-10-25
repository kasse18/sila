package user

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go-template/internal/models/models"
	"go-template/internal/repository"
)

type User struct {
	db *sqlx.DB
}

func (u User) Get(ctx context.Context, id int) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u User) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (u User) Create(ctx context.Context, userCreate models.CreateUser) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (u User) GetPwdByEmail(ctx context.Context, email string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (u User) GetIDByEmail(ctx context.Context, email string) (int, error) {
	//TODO implement me
	panic("implement me")
}

func InitUserRepo(db *sqlx.DB) repository.UserRepo {
	return User{
		db: db,
	}
}
