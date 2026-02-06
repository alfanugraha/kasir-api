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

func (s *ProductService) GetAllProducts(name string) ([]model.Product, error) {
	return s.repo.GetAllProducts(name)
}

func (s *ProductService) Create(input *model.ProductInput) (*model.Product, error) {
	return s.repo.Create(input)
}

func (s *ProductService) GetProductByID(id int) (*model.Product, error) {
	return s.repo.GetProductByID(id)
}

func (s *ProductService) Update(id int, input *model.ProductInput) (*model.Product, error) {
	return s.repo.Update(id, input)
}

func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}
