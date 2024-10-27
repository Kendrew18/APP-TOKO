package cabang

import (
	"APP-TOKO/model/Admin/request"
	"APP-TOKO/service/Admin/cabang"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InputCabang(c echo.Context) error {
	var Request request.Input_Cabang_Request

	err := c.Bind(&Request)

	fmt.Println(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := cabang.Input_Cabang(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadCabang(c echo.Context) error {
	result, err := cabang.Read_Cabang()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
