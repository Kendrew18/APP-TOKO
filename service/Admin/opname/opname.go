package opname

import (
	"APP-TOKO/db"
	"APP-TOKO/model/Admin/request"
	"APP-TOKO/model/Admin/response"
	"net/http"
)

func Start_Opname(Request request.Start_Opname_Request) (response.Response, error) {
	var res response.Response
	var data response.Start_Opname_Response
	var result_check response.Start_Opname_Response_Database

	con := db.CreateConGorm()

	err := con.Table("BARCODE").Select("id_barcode", "status").Where("barcode = ?", Request.Barcode).Scan(&result_check)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	if result_check.Id_barcode != "" && result_check.Status == 0 {
		data.Status = "Barcode Tidak Aktif"
	} else if result_check.Id_barcode != "" && result_check.Status == 1 {
		data.Status = "Barcode Aktif"
	} else {
		data.Status = "Barcode Tidak Ditemukan"
	}

	if data.Status == "" {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = data
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = data
	}

	return res, nil
}
