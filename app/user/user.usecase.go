package user

import (
	"github.com/jihanlugas/pos/app/authentication"
	"github.com/jihanlugas/pos/db"
	"github.com/jihanlugas/pos/model"
	"github.com/jihanlugas/pos/request"
)

type Usecase interface {
	GetById(id string) (model.UserView, error)
	Create(loginUser authentication.UserLogin, req *request.CreateUser) error
}

type userUsecase struct {
	repo Repository
}

func (u userUsecase) Create(loginUser authentication.UserLogin, req *request.CreateUser) error {
	var err error
	var data model.User

	data = model.User{
		RoleID:      "",
		Email:       req.Email,
		Username:    req.Username,
		NoHp:        req.NoHp,
		Fullname:    req.Fullname,
		Passwd:      req.Passwd,
		PassVersion: 1,
		Active:      true,
		PhotoID:     "",
		LastLoginDt: nil,
		CreateBy:    loginUser.UserID,
		UpdateBy:    loginUser.UserID,
	}

	conn, closeConn := db.GetConnection()
	defer closeConn()

	tx := conn.Begin()

	err = u.repo.Create(tx, data)
	if err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

func (u userUsecase) GetById(id string) (model.UserView, error) {
	var err error

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
