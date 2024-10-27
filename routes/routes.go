package routes

import (
	"APP-TOKO/controller/Admin/cabang"
	"APP-TOKO/controller/Admin/provider"
	"APP-TOKO/controller/Admin/tipe"
	"APP-TOKO/controller/Admin/user_app"
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
	US.POST("/login", user_app.Login)
	US.POST("/sign-up", user_app.SignUp)

	//cabang
	CB := e.Group("/CB")
	CB.POST("/cabang", cabang.InputCabang)
	CB.GET("/cabang", cabang.ReadCabang)

	//provider
	P := e.Group("/P")
	P.POST("/provider", provider.InputProvider)
	P.GET("/provider", provider.ReadProvider)

	//tipe
	TP := e.Group("/TP")
	TP.POST("/tipe", tipe.InputTipe)
	TP.GET("/tipe", tipe.ReadTipe)

	return e
}
