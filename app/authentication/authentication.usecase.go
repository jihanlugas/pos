package authentication

import "github.com/labstack/echo/v4"

type Usecase interface {
	Tes(c echo.Context) string
	Test(c echo.Context) string
}

type authenticationUsecase struct {
	repo Repository
}

func (u authenticationUsecase) Tes(c echo.Context) string {
	return "Tes"
}

func (u authenticationUsecase) Test(c echo.Context) string {
	return "Test"
}

func NewAuthenticationUsecase(repo Repository) Usecase {
	return authenticationUsecase{
		repo: repo,
	}
}
