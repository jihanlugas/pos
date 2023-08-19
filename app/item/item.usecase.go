package item

import (
	"github.com/jihanlugas/pos/app/user"
	"github.com/jihanlugas/pos/model"
	"github.com/jihanlugas/pos/request"
)

type Usecase interface {
	GetById(id string) (model.Item, error)
	GetViewById(id string) (model.ItemView, error)
	Create(loginUser user.UserLogin, req *request.CreateItem) error
	Update(loginUser user.UserLogin, id string, req *request.UpdateItem) error
	Delete(loginUser user.UserLogin, id string) error
	Page(req *request.PageUser) ([]model.ItemView, int64, error)
}

type itemUsecase struct {
	repo Repository
}

func (i itemUsecase) GetById(id string) (model.Item, error) {
	//TODO implement me
	panic("implement me")
}

func (i itemUsecase) GetViewById(id string) (model.ItemView, error) {
	//TODO implement me
	panic("implement me")
}

func (i itemUsecase) Create(loginUser user.UserLogin, req *request.CreateItem) error {
	//TODO implement me
	panic("implement me")
}

func (i itemUsecase) Update(loginUser user.UserLogin, id string, req *request.UpdateItem) error {
	//TODO implement me
	panic("implement me")
}

func (i itemUsecase) Delete(loginUser user.UserLogin, id string) error {
	//TODO implement me
	panic("implement me")
}

func (i itemUsecase) Page(req *request.PageUser) ([]model.ItemView, int64, error) {
	//TODO implement me
	panic("implement me")
}

func NewItemUsecase(repo Repository) Usecase {
	return itemUsecase{
		repo: repo,
	}
}
