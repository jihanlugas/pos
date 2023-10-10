package usercompany

import (
	"github.com/jihanlugas/pos/model"
	"github.com/jihanlugas/pos/request"
	"gorm.io/gorm"
)

type Repository interface {
	GetById(conn *gorm.DB, id string) (model.Usercompany, error)
	GetByUserId(conn *gorm.DB, userID string) (model.Usercompany, error)
	GetByCompanyId(conn *gorm.DB, companyID string) (model.Usercompany, error)
	GetCompanyDefaultByUserId(conn *gorm.DB, userID string) (model.Usercompany, error)
	GetViewById(conn *gorm.DB, id string) (model.UsercompanyView, error)
	GetViewByUserId(conn *gorm.DB, userID string) (model.UsercompanyView, error)
	GetViewByCompanyId(conn *gorm.DB, companyID string) (model.UsercompanyView, error)
	GetViewCompanyDefaultByUserId(conn *gorm.DB, userID string) (model.UsercompanyView, error)
	Create(conn *gorm.DB, data model.Usercompany) error
	Update(conn *gorm.DB, data model.Usercompany) error
	Delete(conn *gorm.DB, data model.Usercompany) error
	Page(conn *gorm.DB, req *request.PageUsercompany) ([]model.UsercompanyView, int64, error)
}
