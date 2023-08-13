package user

import (
	"github.com/jihanlugas/pos/model"
	"github.com/jihanlugas/pos/utils"
	"gorm.io/gorm"
	"time"
)

type Repository interface {
	GetById(conn *gorm.DB, id string) (model.User, error)
	GetByUsername(conn *gorm.DB, username string) (model.User, error)
	GetByEmail(conn *gorm.DB, email string) (model.User, error)
	GetByNoHp(conn *gorm.DB, noHp string) (model.User, error)
	GetViewById(conn *gorm.DB, id string) (model.UserView, error)
	GetViewByUsername(conn *gorm.DB, username string) (model.UserView, error)
	GetViewByEmail(conn *gorm.DB, email string) (model.UserView, error)
	GetViewByNoHp(conn *gorm.DB, noHp string) (model.UserView, error)
	Create(conn *gorm.DB, data model.User) error
	Update(conn *gorm.DB, data model.User) error
	Delete(conn *gorm.DB, data model.User) error
	//Update(conn *gorm.DB, id string, data model.User) error
	//Delete(conn *gorm.DB, id string) error
	//Page(conn *gorm.DB) ([]model.UserView, error)
	//List(conn *gorm.DB) ([]model.UserView, error)
}

type userRepository struct {
}

func (u userRepository) GetById(conn *gorm.DB, id string) (model.User, error) {
	var err error
	var data model.User

	err = conn.Where("id = ? ", id).First(&data).Error
	return data, err
}

func (u userRepository) GetByEmail(conn *gorm.DB, email string) (model.User, error) {
	var err error
	var data model.User

	err = conn.Where("email = ? ", email).First(&data).Error
	return data, err
}

func (u userRepository) GetByUsername(conn *gorm.DB, username string) (model.User, error) {
	var err error
	var data model.User

	err = conn.Where("username = ? ", username).First(&data).Error
	return data, err
}

func (u userRepository) GetByNoHp(conn *gorm.DB, noHp string) (model.User, error) {
	var err error
	var data model.User

	err = conn.Where("no_hp = ? ", utils.FormatPhoneTo62(noHp)).First(&data).Error
	return data, err
}

func (u userRepository) GetViewById(conn *gorm.DB, id string) (model.UserView, error) {
	var err error
	var data model.UserView

	err = conn.Where("id = ? ", id).First(&data).Error
	return data, err
}

func (u userRepository) GetViewByEmail(conn *gorm.DB, email string) (model.UserView, error) {
	var err error
	var data model.UserView

	err = conn.Where("email = ? ", email).First(&data).Error
	return data, err
}

func (u userRepository) GetViewByUsername(conn *gorm.DB, username string) (model.UserView, error) {
	var err error
	var data model.UserView

	err = conn.Where("username = ? ", username).First(&data).Error
	return data, err
}

func (u userRepository) GetViewByNoHp(conn *gorm.DB, noHp string) (model.UserView, error) {
	var err error
	var data model.UserView

	err = conn.Where("no_hp = ? ", utils.FormatPhoneTo62(noHp)).First(&data).Error
	return data, err
}

func (u userRepository) Create(conn *gorm.DB, data model.User) error {
	conn.Save(&data)

	return nil
}

func (u userRepository) Update(conn *gorm.DB, data model.User) error {
	conn.Save(&data)

	return nil
}

func (u userRepository) Delete(conn *gorm.DB, data model.User) error {
	now := time.Now()
	data.DeleteDt = &now
	conn.Save(&data)

	return nil
}

func NewUserRepository() Repository {
	return userRepository{}
}
