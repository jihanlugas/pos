package user

import (
	"errors"
	"github.com/jihanlugas/pos/app/app"
	"github.com/jihanlugas/pos/request"
	"github.com/jihanlugas/pos/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type Handler struct {
	usecase Usecase
}

func UserHandler(usecase Usecase) Handler {
	return Handler{
		usecase: usecase,
	}
}

// GetById
// @Tags User
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /user/{id} [get]
func (h Handler) GetById(c echo.Context) error {

	ID := c.Param("id")
	if ID == "" {
		return response.Error(http.StatusBadRequest, "data not found", response.Payload{}).SendJSON(c)
	}

	data, err := h.usecase.GetById(ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.Error(http.StatusBadRequest, "data not found", response.Payload{}).SendJSON(c)
		}
		app.ErrorInternal(c, err)
	}

	res := response.User(data)

	return response.Success(http.StatusOK, "success", res).SendJSON(c)
}

// Create
// @Tags User
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param req body request.CreateUser true "json req body"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /user [post]
func (h Handler) Create(c echo.Context) error {
	var err error

	loginUser, err := GetUserLoginInfo(c)
	if err != nil {
		return response.Error(http.StatusBadRequest, "error validation", response.ValidationError(err)).SendJSON(c)
	}

	req := new(request.CreateUser)
	if err = c.Bind(req); err != nil {
		return response.Error(http.StatusBadRequest, "error validation", response.ValidationError(err)).SendJSON(c)
	}

	if err = c.Validate(req); err != nil {
		return response.Error(http.StatusBadRequest, "error validation", response.ValidationError(err)).SendJSON(c)
	}

	err = h.usecase.Create(loginUser, req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.Error(http.StatusBadRequest, "data not found", response.Payload{}).SendJSON(c)
		}
		app.ErrorInternal(c, err)
	}

	return response.Success(http.StatusOK, "success", response.Payload{}).SendJSON(c)
}
