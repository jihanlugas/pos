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
	GetById(id string) (model.User, error)
	GetViewById(id string) (model.UserView, error)
	Create(loginUser UserLogin, req *request.CreateUser) error
	Update(loginUser UserLogin, id string, req *request.UpdateUser) error
	Delete(loginUser UserLogin, id string) error
	Page(req *request.PageUser) ([]model.UserView, int64, error)
}

type usecaseUser struct {
	repo Repository
}

func (u usecaseUser) Create(loginUser UserLogin, req *request.CreateUser) error {
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

func (u usecaseUser) GetById(id string) (model.User, error) {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err := u.repo.GetById(conn, id)

	return data, err
}

func (u usecaseUser) GetViewById(id string) (model.UserView, error) {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err := u.repo.GetViewById(conn, id)

	return data, err
}

func (u usecaseUser) Update(loginUser UserLogin, id string, req *request.UpdateUser) error {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err := u.repo.GetById(conn, id)
	if err != nil {
		return err
	}

	data.Fullname = req.Fullname
	data.Email = req.Email
	data.Username = req.Username
	data.NoHp = utils.FormatPhoneTo62(req.NoHp)
	data.UpdateBy = loginUser.UserID

	tx := conn.Begin()

	err = u.repo.Update(tx, data)
	if err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

func (u usecaseUser) Delete(loginUser UserLogin, id string) error {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err := u.repo.GetById(conn, id)
	if err != nil {
		return err
	}

	data.DeleteBy = loginUser.UserID

	tx := conn.Begin()

	err = u.repo.Delete(tx, data)
	if err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

func (u usecaseUser) Page(req *request.PageUser) ([]model.UserView, int64, error) {
	var err error
	var data []model.UserView
	var count int64

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, count, err = u.repo.Page(conn, req)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func NewUserUsecase(repo Repository) Usecase {
	return usecaseUser{
		repo: repo,
	}
}
