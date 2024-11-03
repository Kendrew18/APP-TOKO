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

	barcode := ""

	con := db.CreateConGorm().Table("BARCODE")

	err := con.Select("id_barcode").Where("BARCODE = ?", Request.Barcode).Order("co ASC").Scan(&barcode).Error

	if barcode == "" {

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

		Request.Co = co + 1
		Request.Id_barcode = "BR-" + strconv.Itoa(Request.Co)

		date, _ := time.Parse("02-01-2006", Request.Exp_date)
		Request.Exp_date = date.Format("2006-01-02")

		Request.Id_cabang = "CB-0"

		err = con.Select("co", "id_barcode", "id_cabang", "id_tipe", "id_provider", "id_barang", "harga", "barcode", "exp_date").Create(&Request)

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

func Read_Barcode() (response.Response, error) {
	var res response.Response
	var arr_data []response.Read_Barcode_Response

	con := db.CreateConGorm()

	err := con.Table("BARCODE").Select("id_barcode", "nama_cabang", "nama_tipe", "nama_provider", "nama_barang", "harga", "barcode", "status", "DATE_FORMAT(exp_date, '%d-%m-%Y') AS exp_date").Joins("JOIN CABANG ON CABANG.id_cabang = BARCODE.id_cabang").Joins("JOIN TIPE ON TIPE.id_tipe = BARCODE.id_tipe").Joins("JOIN PROVIDER ON PROVIDER.id_provider = BARCODE.id_provider").Joins("JOIN BARANG ON BARANG.id_barang = BARCODE.id_barang").Scan(&arr_data)

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
