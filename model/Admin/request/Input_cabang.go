package request

type Input_Cabang_Request struct {
	Co            int    `json:"co"`
	Id_cabang     string `json:"id_cabang"`
	Nama_Cabang   string `json:"nama_cabang"`
	Alamat_cabang string `json:"alamat_cabang"`
}
