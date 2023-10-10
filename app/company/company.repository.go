package company

import (
	"github.com/jihanlugas/pos/model"
	"github.com/jihanlugas/pos/request"
	"gorm.io/gorm"
)

type Repository interface {
	GetById(conn *gorm.DB, id string) (model.Company, error)
	GetByName(conn *gorm.DB, name string) (model.Company, error)
	GetViewById(conn *gorm.DB, id string) (model.CompanyView, error)
	GetViewByName(conn *gorm.DB, username string) (model.CompanyView, error)
	Create(conn *gorm.DB, data model.Company) error
	Update(conn *gorm.DB, data model.Company) error
	Delete(conn *gorm.DB, data model.Company) error
	Page(conn *gorm.DB, req *request.PageCompany) ([]model.CompanyView, int64, error)
}
