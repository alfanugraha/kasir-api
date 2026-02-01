package repository

import (
	"database/sql"
	"errors"
	"kasir-api/model"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (repo *CategoryRepository) GetAllCategories() ([]model.Category, error) {
	query := "SELECT id, category, description FROM category"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	category := make([]model.Category, 0)
	for rows.Next() {
		var p model.Category
		err := rows.Scan(
			&p.ID,
			&p.Category,
			&p.Description,
		)
		if err != nil {
			return nil, err
		}
		category = append(category, p)
	}
	return category, nil
}

func (repo *CategoryRepository) Create(category *model.Category) error {
	query := "INSERT INTO category (category, description) VALUES ($1, $2) RETURNING id"
	err := repo.db.QueryRow(query, category.Category, category.Description).Scan(&category.ID)
	return err
}

// GetCategoryByID
func (repo *CategoryRepository) GetCategoryByID(id int) (*model.Category, error) {
	query := "SELECT id, category, description FROM category WHERE id = $1"

	var c model.Category
	err := repo.db.QueryRow(query, id).Scan(
		&c.ID,
		&c.Category,
		&c.Description,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("No category found")
	}
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (repo *CategoryRepository) Update(category *model.Category) error {
	query := "UPDATE category SET category = $1, description = $2 WHERE id = $3"
	result, err := repo.db.Exec(query, category.Category, category.Description, category.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("No category found")
	}

	return nil
}

func (repo *CategoryRepository) Delete(id int) error {
	query := "DELETE FROM category WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("No category found")
	}

	return err
}
