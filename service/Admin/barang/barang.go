package barang

import (
	"APP-TOKO/db"
	"APP-TOKO/model/Admin/request"
	"APP-TOKO/model/Admin/response"
	"fmt"
	"net/http"
	"strconv"
)

func Input_Barang(Request request.Input_Barang_Request) (response.Response, error) {
	var res response.Response

	barcode := ""

	con := db.CreateConGorm().Table("BARANG")

	err := con.Select("id_barang").Where("barcode = ?", Request.Barcode).Order("co ASC").Scan(&barcode).Error

	if barcode == "" {

		fmt.Println(Request)

		con := db.CreateConGorm().Table("BARANG")

		co := 0

		err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

		Request.Co = co + 1
		Request.Id_provider = "BR-" + strconv.Itoa(Request.Co)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		Request.Id_cabang = "CB-0"

		err = con.Select("co", "id_barang", "id_cabang", "id_tipe", "id_provider", "nama_barang", "harga", "barcode").Create(&Request)

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

func Read_Barang(Request request.Read_Barang_Request) (response.Response, error) {
	var res response.Response
	var arr_data []response.Read_Barang_Response

	con := db.CreateConGorm()

	err := con.Table("barang").Select("id_barang", "nama_cabang", "nama_tipe", "nama_ptovider", "nama_barang", "harga", "barcode", "status").Joins("JOIN cabang ON cabang.id_cabang = barang.id_cabang").Joins("JOIN tipe ON tipe.id_tipe = barang.id_tipe").Joins("JOIN provider ON provider.id_provider = barang.id_provider").Where("id_cabang = ?", Request.Id_cabang).Scan(&arr_data)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	if arr_data == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_data
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_data
	}

	return res, nil
}
