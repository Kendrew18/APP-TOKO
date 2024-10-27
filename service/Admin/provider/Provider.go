package provider

import (
	"APP-TOKO/db"
	"APP-TOKO/model/Admin/request"
	"APP-TOKO/model/Admin/response"
	"fmt"
	"net/http"
	"strconv"
)

func Input_Provider(Request request.Input_Provider_Request) (response.Response, error) {
	var res response.Response

	id_provider := ""

	con := db.CreateConGorm().Table("PROVIDER")

	err := con.Select("id_provider").Where("nama_cabang = ?", Request.Nama_provider).Order("co ASC").Scan(&id_provider).Error

	if id_provider == "" {

		fmt.Println(Request)

		con := db.CreateConGorm().Table("PROVIDER")

		co := 0

		err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

		Request.Co = co + 1
		Request.Id_provider = "CB-" + strconv.Itoa(Request.Co)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		err = con.Select("co", "id_provider", "nama_provider").Create(&Request)

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
		res.Status = http.StatusNotAcceptable
		res.Message = "cabang telah ada"
		return res, err
	}

	return res, nil
}

func Read_Provider() (response.Response, error) {
	var res response.Response
	var arr_data []response.Read_Provider_Response

	con := db.CreateConGorm().Table("PROVIDER")

	err := con.Select("id_cabang", "nama_cabang", "alamat").Order("co ASC").Scan(&arr_data).Error

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr_data
		return res, err
	}

	if arr_data == nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr_data

	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = arr_data
	}

	return res, nil
}
