package routes

import (
	"APP-TOKO/controller/Admin/barang"
	"APP-TOKO/controller/Admin/barcode"
	"APP-TOKO/controller/Admin/cabang"
	"APP-TOKO/controller/Admin/opname"
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

	AD := e.Group("/AD")

	// user
	US := AD.Group("/US")
	US.POST("/login", user_app.Login)
	US.POST("/sign-up", user_app.SignUp)

	//cabang
	CB := AD.Group("/CB")
	CB.POST("/cabang", cabang.InputCabang)
	CB.GET("/cabang", cabang.ReadCabang)

	//provider
	P := AD.Group("/P")
	P.POST("/provider", provider.InputProvider)
	P.GET("/provider", provider.ReadProvider)

	//tipe
	TP := AD.Group("/TP")
	TP.POST("/tipe", tipe.InputTipe)
	TP.GET("/tipe", tipe.ReadTipe)

	//barcode
	BR := AD.Group("/BR")
	BR.POST("/barcode", barcode.InputBarcode)
	BR.GET("/barcode", barcode.ReadBarcode)

	//barang
	B := AD.Group("/B")
	B.POST("/barang", barang.InputBarang)
	B.GET("/barang", barang.ReadBarang)

	//Opname
	OP := AD.Group("/OP")
	OP.GET("/start-opname", opname.StartOpname)

	return e
}
