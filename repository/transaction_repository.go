package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"kasir-api/model"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (repo *TransactionRepository) Checkout(items []model.CheckoutItem) (*model.Transaction, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	totalPrice := 0.0                             // initiate subtotal -> total all transaction
	details := make([]model.TransactionDetail, 0) // initiate details model -> later insert to db
	// loop items
	for _, item := range items {
		var productPrice float64
		var stock int
		var productName string

		err := tx.QueryRow("SELECT name, price, stock FROM product WHERE id=$1", item.ProductID).Scan(&productName, &productPrice, &stock)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Product with ID %d not found", item.ProductID)
		}
		if err != nil {
			return nil, err
		}

		subtotal := productPrice * float64(item.Quantity)
		totalPrice += subtotal

		_, err = tx.Exec("UPDATE product SET stock = stock - $1 WHERE id = $2 AND stock >= $1", item.Quantity, item.ProductID)
		if err != nil {
			return nil, err
		}

		details = append(details, model.TransactionDetail{
			ProductID:   item.ProductID,
			ProductName: productName,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	var transactionID int
	err = tx.QueryRow("INSERT INTO transaction (total_price) VALUES ($1) RETURNING id", totalPrice).Scan(&transactionID)
	if err != nil {
		return nil, err
	}

	for i := range details {
		_, err = tx.Exec("INSERT INTO transaction_detail (transaction_id, product_id, quantity, subtotal) VALUES ($1, $2, $3, $4)",
			transactionID, details[i].ProductID, details[i].Quantity, details[i].Subtotal)
		if err != nil {
			return nil, err
		}
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &model.Transaction{
		ID:         transactionID,
		TotalPrice: totalPrice,
		Details:    details,
	}, nil
}

func (repo *TransactionRepository) GetTodayTransactions() (*model.TransactionReportRequest, error) {
	var report model.TransactionReportRequest
	err := repo.db.QueryRow(`
			SELECT
				COUNT(*) AS total_transaksi,
				SUM(total_price) AS total_revenue
			FROM transaction
			WHERE created_at::date = CURRENT_DATE`).Scan(
		&report.TotalTransactions,
		&report.TotalRevenue,
	)
	if err != nil {
		return nil, err
	}

	if report.TotalTransactions == 0 {
		return nil, errors.New("Tidak ada transaksi hari ini")
	}

	err = repo.db.QueryRow(`
			SELECT 
				p.name, SUM(td.quantity) AS qty_terjual
			FROM transaction_detail td
			JOIN product p ON td.product_id = p.id
			JOIN transaction t ON td.transaction_id = t.id
			WHERE t.created_at::date = CURRENT_DATE
			GROUP BY p.name
			ORDER BY qty_terjual DESC
			LIMIT 1`).Scan(
		&report.BestSellingProducts.Name,
		&report.BestSellingProducts.Quantity,
	)
	if err != nil {
		return nil, err
	}

	return &report, nil
}

func (repo *TransactionRepository) GetTransactionsByDateRange(startDate string, endDate string) (*model.TransactionReportRequest, error) {
	var report model.TransactionReportRequest

	err := repo.db.QueryRow(`
		SELECT
			COUNT(*) AS total_transaksi, 
			COALESCE(SUM(total_price), 0) AS total_revenue
		FROM transaction
		WHERE created_at::date BETWEEN $1 AND $2`, startDate, endDate).Scan(
		&report.TotalTransactions,
		&report.TotalRevenue,
	)
	if err != nil {
		return nil, err
	}

	if report.TotalTransactions == 0 {
		return nil, errors.New("Tidak ada transaksi di periode waktu tersebut")
	}

	err = repo.db.QueryRow(`
		SELECT 
			p.name, SUM(td.quantity) AS qty_terjual
		FROM transaction_detail td
		JOIN product p ON td.product_id = p.id
		JOIN transaction t ON td.transaction_id = t.id
		WHERE t.created_at::date BETWEEN $1 AND $2
		GROUP BY p.name
		ORDER BY qty_terjual DESC
		LIMIT 1`, startDate, endDate).Scan(
		&report.BestSellingProducts.Name,
		&report.BestSellingProducts.Quantity,
	)
	if err != nil {
		return nil, err
	}

	return &report, nil
}
