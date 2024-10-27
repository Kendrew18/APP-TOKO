package tipe

import (
	"APP-TOKO/model/Admin/request"
	"APP-TOKO/service/Admin/tipe"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InputTipe(c echo.Context) error {
	var Request request.Input_Tipe_Request

	err := c.Bind(&Request)

	fmt.Println(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := tipe.Input_Tipe(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadTipe(c echo.Context) error {
	result, err := tipe.Read_Tipe()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
