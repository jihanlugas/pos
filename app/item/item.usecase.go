package item

import (
	"github.com/jihanlugas/pos/app/user"
	"github.com/jihanlugas/pos/db"
	"github.com/jihanlugas/pos/model"
	"github.com/jihanlugas/pos/request"
)

type Usecase interface {
	GetById(id string) (model.ItemView, error)
	Create(loginUser user.UserLogin, req *request.CreateItem) error
	Update(loginUser user.UserLogin, id string, req *request.UpdateItem) error
	Delete(loginUser user.UserLogin, id string) error
	Page(req *request.PageItem) ([]model.ItemView, int64, error)
}

type usecaseItem struct {
	repo Repository
}

func (u usecaseItem) GetById(id string) (model.ItemView, error) {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err := u.repo.GetViewById(conn, id)

	return data, err
}

func (u usecaseItem) Create(loginUser user.UserLogin, req *request.CreateItem) error {
	var err error
	var data model.Item

	data = model.Item{
		CompanyID:   "",
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
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

func (u usecaseItem) Update(loginUser user.UserLogin, id string, req *request.UpdateItem) error {
	//TODO implement me
	panic("implement me")
}

func (u usecaseItem) Delete(loginUser user.UserLogin, id string) error {
	//TODO implement me
	panic("implement me")
}

func (u usecaseItem) Page(req *request.PageItem) ([]model.ItemView, int64, error) {
	//TODO implement me
	panic("implement me")
}

func NewItemUsecase(repo Repository) Usecase {
	return usecaseItem{
		repo: repo,
	}
}
