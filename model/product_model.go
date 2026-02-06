package model

type Product struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Price    float64  `json:"price"`
	Stock    int      `json:"stock"`
	Category Category `json:"category"`
}

type ProductInput struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Category_ID int     `json:"category_id"`
}
