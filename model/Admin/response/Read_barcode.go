package response

type Read_Barcode_Response struct {
	Id_barcode    string `json:"id_barcode"`
	Nama_cabang   string `json:"nama_cabang"`
	Nama_tipe     string `json:"nama_tipe"`
	Nama_provider string `json:"nama_provider"`
	Nama_barang   string `json:"nama_barang"`
	Harga         int64  `json:"harga"`
	Barcode       string `json:"barcode"`
	Status        string `json:"status"`
}
