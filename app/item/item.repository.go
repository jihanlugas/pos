package item

import (
	"github.com/jihanlugas/pos/model"
	"github.com/jihanlugas/pos/request"
	"gorm.io/gorm"
	"time"
)

type Repository interface {
	GetById(conn *gorm.DB, id string) (model.Item, error)
	GetViewById(conn *gorm.DB, id string) (model.ItemView, error)
	Create(conn *gorm.DB, data model.Item) error
	Update(conn *gorm.DB, data model.Item) error
	Delete(conn *gorm.DB, data model.Item) error
	Page(conn *gorm.DB, req *request.PageItem) ([]model.ItemView, int64, error)
}

type repositoryItem struct {
}

func (r repositoryItem) GetById(conn *gorm.DB, id string) (model.Item, error) {
	var err error
	var data model.Item

	err = conn.Where("id = ? ", id).First(&data).Error
	return data, err
}

func (r repositoryItem) GetViewById(conn *gorm.DB, id string) (model.ItemView, error) {
	var err error
	var data model.ItemView

	err = conn.Where("id = ? ", id).First(&data).Error
	return data, err
}

func (r repositoryItem) Create(conn *gorm.DB, data model.Item) error {
	conn.Create(&data)

	return nil
}

func (r repositoryItem) Update(conn *gorm.DB, data model.Item) error {
	conn.Save(&data)

	return nil
}

func (r repositoryItem) Delete(conn *gorm.DB, data model.Item) error {
	now := time.Now()
	data.DeleteDt = &now
	conn.Save(&data)

	return nil
}

func (r repositoryItem) Page(conn *gorm.DB, req *request.PageItem) ([]model.ItemView, int64, error) {
	var err error
	var data []model.ItemView
	var count int64

	err = conn.Model(&data).
		Where("name LIKE ?", "%"+req.Name+"%").
		Where("description LIKE ?", "%"+req.Description+"%").
		Count(&count).Error
	if err != nil {
		return data, count, err
	}

	err = conn.
		Where("name LIKE ?", "%"+req.Name+"%").
		Where("description LIKE ?", "%"+req.Description+"%").
		Offset((req.GetPage() - 1) * req.GetLimit()).
		Limit(req.GetLimit()).
		Find(&data).Error
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func NewItemRepository() Repository {
	return repositoryItem{}
}
