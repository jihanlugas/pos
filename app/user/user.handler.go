package user

import (
	"errors"
	"github.com/jihanlugas/pos/app/app"
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
