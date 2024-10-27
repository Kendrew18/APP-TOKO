package response

type Login_Response struct {
	Id_user  string `json:"id_user"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Alamat   string `json:"alamat"`
	Token    string `json:"token"`
	Level    int    `json:"level"`
}
