package request

type Sign_Up_Request struct {
	Co        int    `json:"co"`
	Id_User   string `json:"id_user"`
	Id_Cabang string `json:"id_cabang"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Token     string `json:"token"`
	Level     int    `json:"level"`
}
