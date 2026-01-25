package model

type Produk struct {
	ID          int     `json:"id"`
	Nama        string  `json:"nama"`
	Harga       float64 `json:"harga"`
	Stok        int     `json:"stok"`
	ID_Kategori int     `json:"id_category"`
}
