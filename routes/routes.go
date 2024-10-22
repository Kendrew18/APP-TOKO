package routes

import (
	"APP-TOKO/controller/cabang"
	"APP-TOKO/controller/user_app"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Project-APP-Toko")
	})

	// user
	US := e.Group("/US")
	US.POST("/sign_up", user_app.SignUp)

	//cabang
	CB := e.Group("/CB")
	CB.POST("/cabang", cabang.Input_Cabang)

	return e
}
