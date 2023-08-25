package item

import (
	"github.com/jihanlugas/pos/app/user"
	"github.com/jihanlugas/pos/db"
	"github.com/jihanlugas/pos/model"
	"github.com/jihanlugas/pos/request"
)

type Usecase interface {
	GetById(id string) (model.Item, error)
	GetViewById(id string) (model.ItemView, error)
	Create(loginUser user.UserLogin, req *request.CreateItem) error
	Update(loginUser user.UserLogin, id string, req *request.UpdateItem) error
	Delete(loginUser user.UserLogin, id string) error
	Page(req *request.PageItem) ([]model.ItemView, int64, error)
}

type usecaseItem struct {
	repo Repository
}

func (u usecaseItem) GetById(id string) (model.Item, error) {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err := u.repo.GetById(conn, id)

	return data, err
}

func (u usecaseItem) GetViewById(id string) (model.ItemView, error) {
	//TODO implement me
	panic("implement me")
}

func (u usecaseItem) Create(loginUser user.UserLogin, req *request.CreateItem) error {
	//TODO implement me
	panic("implement me")
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
