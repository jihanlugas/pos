package authentication

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	AuthenticationUsecase Usecase
}

func AuthenticationHandler() Handler {
	return Handler{}
}

// SignIn Sign in user
// @Summary Sign in a user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param req body request.Signin true "json req body"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /sign-in [post]
func (h Handler) SignIn(c echo.Context) error {
	//tes := h.AuthenticationUsecase.Tes(c)

	return c.String(http.StatusOK, "tes")
}

// SignOut Sign out user
// @Summary Sign out a user
// @Tags Authentication
// @Accept json
// @Produce json
// // @Param req body request.Signin true "json req body"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /sign-out [post]
func (h Handler) SignOut(c echo.Context) error {
	return nil
}

// SignIn Sign Up user
// @Summary Sign Up a user
// @Tags Authentication
// @Accept json
// @Produce json
// // @Param req body request.Signin true "json req body"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /sign-up [post]
func (h Handler) SignUp(c echo.Context) error {
	return nil
}

// RefreshToken godoc
// @Tags Authentication
// @Summary To do refresh token
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /refresh-token [get]
func (h Handler) RefreshToken(c echo.Context) error {
	return nil
}

func (h Handler) ResetPassword(c echo.Context) error {
	return nil
}
