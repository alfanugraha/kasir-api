package model

type Category struct {
	ID          int    `json:"id"`
	Category    string `json:"category"`
	Description string `json:"description"`
}
