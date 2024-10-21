package user_app

import (
	"APP-TOKO/model/request"
	"APP-TOKO/service/user_app"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SignUp(c echo.Context) error {
	var Request request.Sign_Up_Request

	err := c.Bind(&Request)

	fmt.Println(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := user_app.Sign_Up(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
