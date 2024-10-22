package user_app

import (
	"APP-TOKO/db"
	"APP-TOKO/model/request"
	"APP-TOKO/model/response"
	"fmt"
	"net/http"
	"strconv"
)

func Sign_Up(Request request.Sign_Up_Request) (response.Response, error) {
	var res response.Response

	username := ""

	con := db.CreateConGorm().Table("user")

	err := con.Select("username").Where("username = ? && status = 0", Request.Username).Order("co ASC").Scan(&username).Error

	if username == "" {

		fmt.Println(Request)

		con := db.CreateConGorm().Table("user")

		co := 0

		err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

		Request.Co = co + 1
		Request.Id_User = "US-" + strconv.Itoa(Request.Co)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		Request.Status = 0

		err = con.Select("co", "id_user", "id_cabang", "username", "password", "token", "status").Create(&Request)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		} else {
			res.Status = http.StatusOK
			res.Message = "Suksess"
			res.Data = map[string]int64{
				"rows": err.RowsAffected,
			}
		}
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Username Telah ada"
		return res, err
	}

	return res, nil
}
