package barcode

import (
	"APP-TOKO/db"
	"APP-TOKO/model/Admin/request"
	"APP-TOKO/model/Admin/response"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func Input_Barcode(Request request.Input_Barcode_Request) (response.Response, error) {
	var res response.Response

	con := db.CreateConGorm().Table("BARCODE")

	count := 0

	for i := 0; i < len(Request.Barcode); i++ {
		barcode := ""

		err := con.Select("id_barcode").Where("BARCODE = ?", Request.Barcode[i].Barcode).Order("co ASC").Scan(&barcode).Error

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err
		}

		if barcode == "" {
			count++
		}

	}

	if count == len(Request.Barcode) {

		fmt.Println(Request)

		con := db.CreateConGorm().Table("BARCODE")

		co := 0

		err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		for i := 0; i < len(Request.Barcode); i++ {
			Request.Barcode[i].Co = co + 1 + i
			Request.Barcode[i].Id_barcode = "BR-" + strconv.Itoa(Request.Barcode[i].Co)
			date, _ := time.Parse("02-01-2006", Request.Barcode[i].Exp_date)
			Request.Barcode[i].Exp_date = date.Format("2006-01-02")
			Request.Barcode[i].Id_cabang = "CB-0"
		}

		err = con.Select("co", "id_barcode", "id_cabang", "id_tipe", "id_provider", "id_barang", "harga", "barcode", "exp_date").Create(&Request.Barcode)

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
		res.Message = "Barcode Telah Terdaftar"
		return res, nil
	}

	return res, nil
}

func Read_Barcode() (response.Response, error) {
	var res response.Response
	var arr_data []response.Read_Barcode_Response

	con := db.CreateConGorm()

	err := con.Table("BARCODE").Select("id_barcode", "nama_cabang", "nama_tipe", "nama_provider", "nama_barang", "harga", "barcode", "status", "DATE_FORMAT(exp_date, '%d-%m-%Y') AS exp_date").Joins("JOIN CABANG ON CABANG.id_cabang = BARCODE.id_cabang").Joins("JOIN BARANG ON BARANG.id_barang = BARCODE.id_barang").Joins("JOIN TIPE ON TIPE.id_tipe = BARANG.id_tipe").Joins("JOIN PROVIDER ON PROVIDER.id_provider = BARANG.id_provider").Scan(&arr_data)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr_data
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
