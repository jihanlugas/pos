package user

import (
	"github.com/jihanlugas/pos/model"
	"gorm.io/gorm"
)

type Repository interface {
	GetById(conn *gorm.DB, id string) (model.UserView, error)
	GetByUsername(conn *gorm.DB, username string) (model.UserView, error)
	GetByEmail(conn *gorm.DB, email string) (model.UserView, error)
	Create(conn *gorm.DB, data model.User) error
	//Update(conn *gorm.DB, id string, data model.User) error
	//Delete(conn *gorm.DB, id string) error
	//Page(conn *gorm.DB) ([]model.UserView, error)
	//List(conn *gorm.DB) ([]model.UserView, error)
}

type userRepository struct {
}

func (u userRepository) GetById(conn *gorm.DB, id string) (model.UserView, error) {
	var err error
	var data model.UserView

	err = conn.Where("id = ? ", id).First(&data).Error
	return data, err
}

func (u userRepository) GetByEmail(conn *gorm.DB, email string) (model.UserView, error) {
	var err error
	var data model.UserView

	err = conn.Where("email = ? ", email).First(&data).Error
	return data, err
}

func (u userRepository) GetByUsername(conn *gorm.DB, username string) (model.UserView, error) {
	var err error
	var data model.UserView

	err = conn.Where("username = ? ", username).First(&data).Error
	return data, err
}

func (u userRepository) Create(conn *gorm.DB, data model.User) error {
	conn.Save(&data)

	return nil
}

//func (u userRepository) Page(conn *gorm.DB) ([]model.UserView, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (u userRepository) List(conn *gorm.DB) ([]model.UserView, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (u userRepository) Update(conn *gorm.DB, id string, data model.User) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (u userRepository) Delete(conn *gorm.DB, id string) error {
//	//TODO implement me
//	panic("implement me")
//}

func NewUserRepository() Repository {
	return userRepository{}
}
