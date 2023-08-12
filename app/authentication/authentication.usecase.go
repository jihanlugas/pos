package authentication

import (
	"errors"
	"github.com/jihanlugas/pos/app/user"
	"github.com/jihanlugas/pos/cryption"
	"github.com/jihanlugas/pos/db"
	"github.com/jihanlugas/pos/model"
	"github.com/jihanlugas/pos/request"
	"github.com/jihanlugas/pos/utils"
	"github.com/labstack/echo/v4"
)

type Usecase interface {
	Tes(c echo.Context) string
	Test(c echo.Context) string
	SignIn(req *request.Signin) error
}

type authenticationUsecase struct {
	repo     Repository
	userRepo user.Repository
}

func (u authenticationUsecase) SignIn(req *request.Signin) error {
	var err error
	var user model.UserView

	conn, closeConn := db.GetConnection()
	defer closeConn()

	if utils.IsValidEmail(req.Username) {
		user, err = u.userRepo.GetByEmail(conn, req.Username)
	} else {
		user, err = u.userRepo.GetByUsername(conn, req.Username)
	}

	if err != nil {
		return err
	}

	err = cryption.CheckAES64(req.Passwd, user.Passwd)
	if err != nil {
		return errors.New("invalid username or password")
	}

	if !user.Active {
		return errors.New("user not active")
	}

	return err
}

func (u authenticationUsecase) Tes(c echo.Context) string {
	return "Tes"
}

func (u authenticationUsecase) Test(c echo.Context) string {
	return "Test"
}

func NewAuthenticationUsecase(repo Repository, userRepo user.Repository) Usecase {
	return authenticationUsecase{
		repo:     repo,
		userRepo: userRepo,
	}
}
