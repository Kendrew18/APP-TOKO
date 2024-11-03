package request

type Input_Barcode_Request struct {
	Co          int    `json:"co"`
	Id_barcode  string `json:"id_barcode"`
	Id_cabang   string `json:"id_cabang"`
	Id_tipe     string `json:"id_tipe"`
	Id_provider string `json:"id_provider"`
	Id_barang   string `json:"id_barang"`
	Harga       int64  `json:"harga"`
	Barcode     string `json:"barcode"`
	Exp_date    string `json:"exp_date"`
}
