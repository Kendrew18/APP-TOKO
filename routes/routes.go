package routes

import (
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

	TMP := e.Group("/US")

	//NDL
	TMP.GET("/sign_up", user_app.SignUp)

	return e
}
