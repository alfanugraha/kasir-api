package service

import (
	"kasir-api/model"
	"kasir-api/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAllProducts() ([]model.Product, error) {
	return s.repo.GetAllProducts()
}

func (s *ProductService) Create(product *model.Product) error {
	return s.repo.Create(product)
}

func (s *ProductService) GetProductByID(id int) (*model.Product, error) {
	return s.repo.GetProductByID(id)
}

func (s *ProductService) Update(product *model.Product) error {
	return s.repo.Update(product)
}

func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}
