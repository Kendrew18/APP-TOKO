package user_app

import (
	"APP-TOKO/db"
	"APP-TOKO/model/response"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func Sign_Up(Request request.Sign_Up_Request) (response.Response, error) {
	var res response.Response

	username := ""

	con := db.CreateConGorm().Table("user")

	err := con.Select("username").Where("username = ? && email_bisnis = ? && status = 0", Request.Username, Request.Email_bisnis).Order("co ASC").Scan(&username).Error

	if username == "" {

		fmt.Println(Request)

		con := db.CreateConGorm().Table("user")

		co := 0

		err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

		Request.Co = co + 1
		Request.Kode_user = "US-" + strconv.Itoa(Request.Co)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		date, _ := time.Parse("02-01-2006", Request.Birth_date)
		Request.Birth_date = date.Format("2006-01-02")
		Request.Status = -1

		Request.Key = uuid.NewString()

		err = con.Select("co", "kode_user", "nama_lengkap", "birth_date", "gender", "category_bisnis", "nama_bisnis", "alamat_bisnis", "telepon_bisnis", "email_bisnis", "instagram", "facebook", "username", "password", "status").Create(&Request)

		if err.Error != nil {
			fmt.Println("masuk")
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		var Request_OTP request.Otp_Request

		Request_OTP.Kode_otp = Send_Email(Request.Email_bisnis, Request.Nama_lengkap, "noreply", "Softipos")
		Request_OTP.Kode_user = Request.Kode_user

		Request_OTP.Email = Request.Email_bisnis

		layoutFormat := "2006-01-02 15:04:05"

		date_sent_time := time.Now()

		date_sent := date_sent_time.Format(layoutFormat)

		// date_resent_time, _ := time.Parse(layoutFormat, date_sent)

		// date_resent_time = date_resent_time.Add(time.Minute * time.Duration(1))

		// date_resent := date_resent_time.Format(layoutFormat)

		Request_OTP.Time_sent = date_sent
		Request_OTP.Nama_lengkap = Request.Nama_lengkap
		//Request_OTP.Time_resent = date_resent

		fmt.Println(Request_OTP)

		con2 := db.CreateConGorm().Table("otp")

		err = con2.Select("kode_user", "nama_lengkap", "email", "kode_otp", "time_sent").Create(&Request_OTP)

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
