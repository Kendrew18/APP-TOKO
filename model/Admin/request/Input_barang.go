package request

type Input_Barang_Request struct {
	Co          int    `json:"co"`
	Id_tipe     string `json:"id_tipe"`
	Id_provider string `json:"id_provider"`
	Id_barang   string `json:"id_barang"`
	Nama_barang string `json:"nama_barang"`
}
