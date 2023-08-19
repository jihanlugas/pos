package item

import (
	"github.com/jihanlugas/pos/model"
	"github.com/jihanlugas/pos/request"
	"gorm.io/gorm"
)

type Repository interface {
	GetById(conn *gorm.DB, id string) (model.Item, error)
	GetViewById(conn *gorm.DB, id string) (model.ItemView, error)
	Create(conn *gorm.DB, data model.Item) error
	Update(conn *gorm.DB, data model.Item) error
	Delete(conn *gorm.DB, data model.Item) error
	Page(conn *gorm.DB, req *request.PageItem) ([]model.ItemView, int64, error)
}

type itemRepository struct {
}

func (i itemRepository) GetById(conn *gorm.DB, id string) (model.Item, error) {
	//TODO implement me
	panic("implement me")
}

func (i itemRepository) GetViewById(conn *gorm.DB, id string) (model.ItemView, error) {
	//TODO implement me
	panic("implement me")
}

func (i itemRepository) Create(conn *gorm.DB, data model.Item) error {
	//TODO implement me
	panic("implement me")
}

func (i itemRepository) Update(conn *gorm.DB, data model.Item) error {
	//TODO implement me
	panic("implement me")
}

func (i itemRepository) Delete(conn *gorm.DB, data model.Item) error {
	//TODO implement me
	panic("implement me")
}

func (i itemRepository) Page(conn *gorm.DB, req *request.PageItem) ([]model.ItemView, int64, error) {
	//TODO implement me
	panic("implement me")
}

func NewItemRepository() Repository {
	return itemRepository{}
}
