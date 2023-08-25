package user

import (
	"errors"
	"github.com/jihanlugas/pos/config"
	"github.com/jihanlugas/pos/cryption"
	"github.com/jihanlugas/pos/db"
	"github.com/jihanlugas/pos/model"
	"github.com/jihanlugas/pos/request"
	"github.com/jihanlugas/pos/utils"
	"github.com/labstack/echo/v4"
	"time"
)

type AuthenticationUsecase interface {
	Tes(c echo.Context) string
	Test(c echo.Context) string
	SignIn(req *request.Signin) (string, error)
}

type usecaseAuthentication struct {
	repo     AuthenticationRepository
	userRepo Repository
}

func (u usecaseAuthentication) SignIn(req *request.Signin) (string, error) {
	var err error
	var data model.User
	var userLogin UserLogin

	conn, closeConn := db.GetConnection()
	defer closeConn()

	if utils.IsValidEmail(req.Username) {
		data, err = u.userRepo.GetByEmail(conn, req.Username)
	} else {
		data, err = u.userRepo.GetByUsername(conn, req.Username)
	}

	if err != nil {
		return "", err
	}

	err = cryption.CheckAES64(req.Passwd, data.Passwd)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	if !data.Active {
		return "", errors.New("user not active")
	}

	now := time.Now()
	tx := conn.Begin()

	data.LastLoginDt = &now
	data.UpdateBy = data.ID
	err = u.userRepo.Update(tx, data)
	if err != nil {
		return "", err
	}

	err = tx.Commit().Error
	if err != nil {
		return "", err
	}

	expiredAt := time.Now().Add(time.Hour * time.Duration(config.AuthTokenExpiredHour))

	userLogin.UserID = data.ID
	userLogin.RoleID = data.RoleID
	userLogin.PassVersion = data.PassVersion
	token, err := CreateToken(userLogin, expiredAt)
	if err != nil {
		return "", err
	}

	return token, err
}

func (u usecaseAuthentication) Tes(c echo.Context) string {
	return "Tes"
}

func (u usecaseAuthentication) Test(c echo.Context) string {
	return "Test"
}

func NewAuthenticationUsecase(repo AuthenticationRepository, userRepo Repository) AuthenticationUsecase {
	return usecaseAuthentication{
		repo:     repo,
		userRepo: userRepo,
	}
}
