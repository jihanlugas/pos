package authentication

import (
	"github.com/jihanlugas/pos/request"
	"github.com/jihanlugas/pos/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	usecase Usecase
}

func AuthenticationHandler(usecase Usecase) Handler {
	return Handler{
		usecase: usecase,
	}
}

// SignIn Sign in user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param req body request.Signin true "json req body"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /sign-in [post]
func (h Handler) SignIn(c echo.Context) error {
	var err error

	req := new(request.Signin)
	if err = c.Bind(req); err != nil {
		return response.Error(http.StatusBadRequest, "error validation", response.ValidationError(err)).SendJSON(c)
	}

	if err = c.Validate(req); err != nil {
		return response.Error(http.StatusBadRequest, "error validation", response.ValidationError(err)).SendJSON(c)
	}

	err = h.usecase.SignIn(req)
	return response.Success(http.StatusOK, "success", response.Payload{}).SendJSON(c)
}

// SignOut Sign out user
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

// SignUp
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

// RefreshToken
// @Tags Authentication
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
