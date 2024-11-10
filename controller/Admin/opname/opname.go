package opname

import (
	"APP-TOKO/model/Admin/request"
	"APP-TOKO/model/Admin/response"
	"APP-TOKO/service/Admin/opname"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func StartOpname(c echo.Context) error {
	var Request request.Start_Opname_Request
	var result response.Response
	var err error

	Request.Barcode = c.Request().Header.Get("barcode")

	fmt.Println(Request.Barcode)

	result, err = opname.Start_Opname(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
