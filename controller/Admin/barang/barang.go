package barang

import (
	"APP-TOKO/model/Admin/request"
	"APP-TOKO/service/Admin/barang"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InputBarang(c echo.Context) error {
	var Request request.Input_Barang_Request

	err := c.Bind(&Request)

	fmt.Println(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := barang.Input_Barang(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadBarang(c echo.Context) error {
	result, err := barang.Read_Barang()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
