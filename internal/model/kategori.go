package model

type Kategori struct {
	ID         int    `json:"id"`
	Kategori   string `json:"nama"`
	Keterangan string `json:"description"`
}
