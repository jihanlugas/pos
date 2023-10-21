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
	"log"
	"net/http"
)

func Init() *echo.Echo {
	router := websiteRouter()

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

	router.Use(loggerMiddleware)

	//router.Use(logMiddleware)

	router.GET("/swg/*", echoSwagger.WrapHandler)
	router.GET("/", app.Ping)

	router.POST("/sign-in", authenticationHandler.SignIn)
	router.GET("/sign-out", authenticationHandler.SignOut)
	router.POST("/sign-up", authenticationHandler.SignUp)
	router.GET("/refresh-token", authenticationHandler.RefreshToken, checkTokenMiddleware)
	router.GET("/reset-password", authenticationHandler.ResetPassword)

	userRouter := router.Group("/user")
	userRouter.GET("/:id", userHandler.GetById)
	userRouter.POST("", userHandler.Create, checkTokenMiddleware)
	userRouter.PUT("/:id", userHandler.Update, checkTokenMiddleware)
	userRouter.DELETE("/:id", userHandler.Delete, checkTokenMiddleware)
	userRouter.GET("/page", userHandler.Page, checkTokenMiddleware)

	itemRouter := router.Group("/item")
	itemRouter.GET("/:id", itemHandler.GetById)
	itemRouter.POST("", itemHandler.Create, checkTokenMiddleware)
	itemRouter.PUT("/:id", itemHandler.Update, checkTokenMiddleware)
	itemRouter.DELETE("/:id", itemHandler.Delete, checkTokenMiddleware)
	itemRouter.GET("/page", itemHandler.Page, checkTokenMiddleware)

	return router

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

func logMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		body, _ := io.ReadAll(c.Request().Body)
		c.Set(constant.Request, string(body))
		c.Request().Body = io.NopCloser(bytes.NewBuffer(body))

		return next(c)
	}
}

func loggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Log incoming request
		log.Printf("Incoming Request: %s %s", c.Request().Method, c.Request().URL.String())
		fmt.Println("Incoming Request: ", c.Request().Body)

		body, _ := io.ReadAll(c.Request().Body)
		c.Set(constant.Request, string(body))
		c.Request().Body = io.NopCloser(bytes.NewBuffer(body))

		// Call next handler
		if err := next(c); err != nil {
			c.Error(err)
		}

		// Log outgoing response
		fmt.Println("Method: ", c.Request().Method)
		fmt.Println("Path: ", c.Request().URL.Host)
		fmt.Println("Path: ", c.Request().URL.Opaque)
		fmt.Println("Path: ", c.Request().URL.Path)
		fmt.Println("Path: ", c.Request().RequestURI)
		fmt.Println("Path: ", c.Request().URL.RawQuery)
		fmt.Println("Path: ", c.Request().URL.RawPath)
		fmt.Println("Path: ", c.Request().URL.RawFragment)
		fmt.Println("Path: ", c.Request().URL.EscapedPath())
		fmt.Println("Path: ", c.Request().Host)
		fmt.Println("Path: ", c.Request().URL.RequestURI())
		fmt.Println("Path: ", c.Request().URL.String())
		fmt.Println("Request: ", string(body))
		fmt.Println("Response: ", string(c.Get(constant.Response).([]byte)))

		return nil
	}
}

func checkTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
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
