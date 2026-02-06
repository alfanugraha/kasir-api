package model

import "time"

type Transaction struct {
	ID         int                 `json:"id"`
	TotalPrice float64             `json:"total_price"`
	CreatedAt  time.Time           `json:"created_at"`
	Details    []TransactionDetail `json:"details"`
}

type TransactionDetail struct {
	ID            int     `json:"id"`
	TransactionID int     `json:"transaction_id"`
	ProductID     int     `json:"product_id"`
	ProductName   string  `json:"product_name,omitempty"`
	Quantity      int     `json:"quantity"`
	Subtotal      float64 `json:"subtotal"`
}

type CheckoutItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CheckoutRequest struct {
	Items []CheckoutItem `json:"items"`
}

type TransactionReportRequest struct {
	TotalRevenue        float64 `json:"total_revenue"`
	TotalTransactions   int     `json:"total_transaksi"`
	BestSellingProducts struct {
		Name     string `json:"nama"`
		Quantity int    `json:"qty_terjual"`
	} `json:"produk_terlaris"`
}
