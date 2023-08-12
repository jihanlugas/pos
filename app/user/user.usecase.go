package user

import (
	"errors"
	"github.com/jihanlugas/pos/cryption"
	"github.com/jihanlugas/pos/db"
	"github.com/jihanlugas/pos/model"
	"github.com/jihanlugas/pos/request"
	"github.com/jihanlugas/pos/utils"
)

type Usecase interface {
	GetById(id string) (model.UserView, error)
	Create(loginUser UserLogin, req *request.CreateUser) error
}

type userUsecase struct {
	repo Repository
}

func (u userUsecase) Create(loginUser UserLogin, req *request.CreateUser) error {
	var err error
	var data model.User

	password, err := cryption.EncryptAES64("123456")
	if err != nil {
		return errors.New("failed to encrypt")
	}

	data = model.User{
		RoleID:      "",
		Email:       req.Email,
		Username:    req.Username,
		NoHp:        utils.FormatPhoneTo62(req.NoHp),
		Fullname:    req.Fullname,
		Passwd:      password,
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

	data, err := u.repo.GetViewById(conn, id)

	return data, err
}

func NewUserUsecase(repo Repository) Usecase {
	return userUsecase{
		repo: repo,
	}
}
