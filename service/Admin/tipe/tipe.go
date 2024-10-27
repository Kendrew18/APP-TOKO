package tipe

import (
	"APP-TOKO/db"
	"APP-TOKO/model/Admin/request"
	"APP-TOKO/model/Admin/response"
	"fmt"
	"net/http"
	"strconv"
)

func Input_Tipe(Request request.Input_Tipe_Request) (response.Response, error) {
	var res response.Response

	id_cabang := ""

	con := db.CreateConGorm().Table("TIPE")

	err := con.Select("id_tipe").Where("nama_tipe = ?", Request.Nama_tipe).Order("co ASC").Scan(&id_cabang).Error

	if id_cabang == "" {

		fmt.Println(Request)

		con := db.CreateConGorm().Table("TIPE")

		co := 0

		err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

		Request.Co = co + 1
		Request.Id_tipe = "TP-" + strconv.Itoa(Request.Co)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		err = con.Select("co", "id_tipe", "nama_tipe").Create(&Request)

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

func Read_Tipe() (response.Response, error) {
	var res response.Response
	var arr_invent []response.Read_Tipe_Response

	con := db.CreateConGorm().Table("TIPE")

	err := con.Select("id_tipe", "nama_tipe").Order("co ASC").Scan(&arr_invent).Error

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr_invent
		return res, err
	}

	if arr_invent == nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr_invent

	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = arr_invent
	}

	return res, nil
}
