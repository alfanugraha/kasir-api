package service

import (
	"kasir-api/model"
	"kasir-api/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAllCategories() ([]model.Category, error) {
	return s.repo.GetAllCategories()
}

func (s *CategoryService) Create(category *model.Category) error {
	return s.repo.Create(category)
}

func (s *CategoryService) GetCategoryByID(id int) (*model.Category, error) {
	return s.repo.GetCategoryByID(id)
}

func (s *CategoryService) Update(category *model.Category) error {
	return s.repo.Update(category)
}

func (s *CategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}
