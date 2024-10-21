package request

type Sign_Up_Request struct {
	Co        int    `json:"co"`
	Id_User   string `json:"id_user"`
	Id_Cabang string `json:"id_cabang"`
	Username  string `json:"usernama"`
	Password  string `json:"password"`
	Token     string `json:"token"`
	Status    int    `json:"status"`
}
