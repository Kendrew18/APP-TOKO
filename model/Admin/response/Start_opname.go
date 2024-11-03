package response

type Start_Opname_Response struct {
	Status string `json:"status"`
}

type Start_Opname_Response_Database struct {
	Id_barcode string `json:"id_barcode"`
	Status     int    `json:"status"`
}
