package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jihanlugas/pos/app/app"
	"github.com/jihanlugas/pos/app/item"
	"github.com/jihanlugas/pos/app/user"
	"github.com/jihanlugas/pos/config"
	"github.com/jihanlugas/pos/constant"
	"github.com/jihanlugas/pos/db"
	_ "github.com/jihanlugas/pos/docs"
	"github.com/jihanlugas/pos/model"
	"github.com/jihanlugas/pos/response"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"io"
	"net/http"
)

func Init() *echo.Echo {
	r := websiteRouter()
	checkToken := checkTokenMiddleware()
	log := logMiddleware()

	//userController := controller.UserComposer()
	//
	//router.GET("/swg/*", echoSwagger.WrapHandler)
	//
	//router.GET("/", controller.Ping)
	//router.POST("/sign-in", userController.SignIn)
	//router.GET("/sign-out", userController.SignOut)
	//router.GET("/refresh-token", userController.RefreshToken, checkToken)

	authenticationRepo := user.NewAuthenticationRepository()
	userRepo := user.NewUserRepository()
	itemRepo := item.NewItemRepository()

	authenticationUsecase := user.NewAuthenticationUsecase(authenticationRepo, userRepo)
	userUsecase := user.NewUserUsecase(userRepo)
	itemUsecase := item.NewItemUsecase(itemRepo)

	authenticationHandler := user.NewAuthenticationHandler(authenticationUsecase)
	userHandler := user.UserHandler(userUsecase)
	itemHandler := item.ItemHandler(itemUsecase)

	router := r.Group("", log)

	router.GET("/swg/*", echoSwagger.WrapHandler)
	router.GET("/", app.Ping)

	router.POST("/sign-in", authenticationHandler.SignIn)
	router.GET("/sign-out", authenticationHandler.SignOut)
	router.POST("/sign-up", authenticationHandler.SignUp)
	router.GET("/refresh-token", authenticationHandler.RefreshToken, checkToken)
	router.GET("/reset-password", authenticationHandler.ResetPassword)

	userRouter := router.Group("/user")
	userRouter.GET("/:id", userHandler.GetById)
	userRouter.POST("", userHandler.Create, checkToken)
	userRouter.PUT("/:id", userHandler.Update, checkToken)
	userRouter.DELETE("/:id", userHandler.Delete, checkToken)
	userRouter.GET("/page", userHandler.Page, checkToken)

	itemRouter := router.Group("/item")
	itemRouter.GET("/:id", itemHandler.GetById)
	itemRouter.POST("", itemHandler.Create, checkToken)
	itemRouter.PUT("/:id", itemHandler.Update, checkToken)
	itemRouter.DELETE("/:id", itemHandler.Delete, checkToken)
	itemRouter.GET("/page", itemHandler.Page, checkToken)

	return r

}

func httpErrorHandler(err error, c echo.Context) {
	var errorResponse *response.Response
	code := http.StatusInternalServerError
	switch e := err.(type) {
	case *echo.HTTPError:
		// Handle pada saat URL yang di request tidak ada. atau ada kesalahan server.
		code = e.Code
		errorResponse = &response.Response{
			Status:  false,
			Message: fmt.Sprintf("%v", e.Message),
			Payload: map[string]interface{}{},
			Code:    code,
		}
	case *response.Response:
		errorResponse = e
	default:
		// Handle error dari panic
		code = http.StatusInternalServerError
		if config.Debug {
			errorResponse = &response.Response{
				Status:  false,
				Message: err.Error(),
				Payload: map[string]interface{}{},
				Code:    http.StatusInternalServerError,
			}
		} else {
			errorResponse = &response.Response{
				Status:  false,
				Message: "Internal server error",
				Payload: map[string]interface{}{},
				Code:    http.StatusInternalServerError,
			}
		}
	}

	js, err := json.Marshal(errorResponse)
	if err == nil {
		_ = c.Blob(code, echo.MIMEApplicationJSONCharsetUTF8, js)
	} else {
		b := []byte("{error: true, message: \"unresolved error\"}")
		_ = c.Blob(code, echo.MIMEApplicationJSONCharsetUTF8, b)
	}
}

func logMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			body, _ := io.ReadAll(c.Request().Body)
			c.Set(constant.RequestBodyContext, string(body))
			c.Request().Body = io.NopCloser(bytes.NewBuffer(body))

			return next(c)
		}
	}
}

func checkTokenMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var err error

			userLogin, err := user.ExtractClaims(c.Request().Header.Get(config.HeaderAuthName))
			if err != nil {
				return response.ErrorForce(http.StatusUnauthorized, err.Error(), response.Payload{}).SendJSON(c)
			}

			conn, closeConn := db.GetConnection()
			defer closeConn()

			var user model.User
			err = conn.Where("id = ? ", userLogin.UserID).First(&user).Error
			if err != nil {
				return response.ErrorForce(http.StatusUnauthorized, "Token Expired!", response.Payload{}).SendJSON(c)
			}

			if user.PassVersion != userLogin.PassVersion {
				return response.ErrorForce(http.StatusUnauthorized, "Token Expired~", response.Payload{}).SendJSON(c)
			}

			c.Set(constant.TokenUserContext, userLogin)
			return next(c)
		}
	}
}
