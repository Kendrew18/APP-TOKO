package request

type Input_Barang_Request struct {
	Co          int    `json:"co"`
	Id_barang   string `json:"id_barang"`
	Id_cabang   string `json:"id_cabang"`
	Id_tipe     string `json:"id_tipe"`
	Id_provider string `json:"id_provider"`
	Nama_barang string `json:"nama_barang"`
	Harga       int64  `json:"harga"`
	Barcode     string `json:"barcode"`
}
