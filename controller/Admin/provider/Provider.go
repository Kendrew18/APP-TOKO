package provider

import (
	"APP-TOKO/model/Admin/request"
	"APP-TOKO/service/Admin/provider"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InputProvider(c echo.Context) error {
	var Request request.Input_Provider_Request

	err := c.Bind(&Request)

	fmt.Println(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := provider.Input_Provider(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadProvider(c echo.Context) error {
	result, err := provider.Read_Provider()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
