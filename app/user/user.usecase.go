package user

import (
	"github.com/jihanlugas/pos/db"
	"github.com/jihanlugas/pos/model"
)

type Usecase interface {
	GetById(id string) (model.UserView, error)
}

type userUsecase struct {
	repo Repository
}

func (u userUsecase) GetById(id string) (model.UserView, error) {
	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err := u.repo.GetById(conn, id)

	return data, err
}

func NewUserUsecase(repo Repository) Usecase {
	return userUsecase{
		repo: repo,
	}
}
