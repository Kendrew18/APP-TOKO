package cabang

import (
	"APP-TOKO/db"
	"APP-TOKO/model/Admin/request"
	"APP-TOKO/model/Admin/response"
	"fmt"
	"net/http"
	"strconv"
)

func Input_Cabang(Request request.Input_Cabang_Request) (response.Response, error) {
	var res response.Response

	id_cabang := ""

	con := db.CreateConGorm().Table("CABANG")

	err := con.Select("id_cabang").Where("nama_cabang = ? && alamat_cabang = ?", Request.Nama_Cabang, Request.Alamat_cabang).Order("co ASC").Scan(&id_cabang).Error

	if id_cabang == "" {

		fmt.Println(Request)

		con := db.CreateConGorm().Table("CABANG")

		co := 0

		err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

		Request.Co = co + 1
		Request.Id_cabang = "CB-" + strconv.Itoa(Request.Co)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		err = con.Select("co", "id_cabang", "nama_cabang", "alamat_cabang").Create(&Request)

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

func Read_Cabang() (response.Response, error) {
	var res response.Response
	var arr_invent []response.Read_Cabang_Response

	con := db.CreateConGorm().Table("CABANG")

	err := con.Select("id_cabang", "nama_cabang", "alamat").Order("co ASC").Scan(&arr_invent).Error

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
