package service

import (
	"kasir-api/model"
	"kasir-api/repository"
)

type TransactionService struct {
	repo *repository.TransactionRepository
}

func NewTransactionService(repo *repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Checkout(items []model.CheckoutItem) (*model.Transaction, error) {
	return s.repo.Checkout(items)
}

func (s *TransactionService) GetTodayTransactions() (*model.TransactionReportRequest, error) {
	return s.repo.GetTodayTransactions()
}

func (s *TransactionService) GetTransactionsByDateRange(startDate string, endDate string) (*model.TransactionReportRequest, error) {
	return s.repo.GetTransactionsByDateRange(startDate, endDate)
}
