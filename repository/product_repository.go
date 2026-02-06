package repository

import (
	"database/sql"
	"errors"
	"kasir-api/model"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetAllProducts(name string) ([]model.Product, error) {
	args := []interface{}{}
	query := `
		SELECT 
			p.id, p.name, p.price, p.stock, c.id, c.category, c.description
		FROM product p
		JOIN category c ON p.category_id = c.id`

	if name != "" {
		query += " WHERE p.name ILIKE $1"
		args = append(args, "%"+name+"%")
	}

	rows, err := repo.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]model.Product, 0)
	for rows.Next() {
		var p model.Product
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Price,
			&p.Stock,
			&p.Category.ID,
			&p.Category.Category,
			&p.Category.Description,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if len(products) == 0 {
		return nil, errors.New("No products found")
	}
	return products, nil
}

func (repo *ProductRepository) Create(input *model.ProductInput) (*model.Product, error) {
	var productID int
	query := "INSERT INTO product (name, price, stock, category_id) VALUES ($1, $2, $3, $4) RETURNING id"
	err := repo.db.QueryRow(query, input.Name, input.Price, input.Stock, input.Category_ID).Scan(&productID)
	if err != nil {
		return nil, err
	}

	// Fetch the complete product with category information
	return repo.GetProductByID(productID)
}

// GetProductByID
func (repo *ProductRepository) GetProductByID(id int) (*model.Product, error) {
	query := `
		SELECT 
			p.id, p.name, p.price, p.stock, c.id, c.category, c.description
		FROM product p
		JOIN category c ON p.category_id = c.id
		WHERE p.id = $1`

	var p model.Product
	err := repo.db.QueryRow(query, id).Scan(
		&p.ID,
		&p.Name,
		&p.Price,
		&p.Stock,
		&p.Category.ID,
		&p.Category.Category,
		&p.Category.Description,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("No product found")
	}
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (repo *ProductRepository) Update(id int, input *model.ProductInput) (*model.Product, error) {
	query := "UPDATE product SET name = $1, price = $2, stock = $3, category_id = $4 WHERE id = $5"
	result, err := repo.db.Exec(query, input.Name, input.Price, input.Stock, input.Category_ID, id)
	if err != nil {
		return nil, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rows == 0 {
		return nil, errors.New("No product found")
	}

	// Fetch the complete updated product with category information
	return repo.GetProductByID(id)
}

func (repo *ProductRepository) Delete(id int) error {
	query := "DELETE FROM product WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("No product found")
	}

	return err
}
