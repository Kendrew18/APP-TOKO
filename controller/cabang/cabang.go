package cabang

import (
	"APP-TOKO/model/request"
	"APP-TOKO/service/cabang"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Input_Cabang(c echo.Context) error {
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
