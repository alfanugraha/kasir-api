package handler

import (
	"encoding/json"
	"kasir-api/model"
	"kasir-api/service"
	"net/http"
)

type TransactionHandler struct {
	service *service.TransactionService
}

func NewTransactionHandler(service *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (h *TransactionHandler) HandleCheckout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Checkout(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TransactionHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	var request model.CheckoutRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	transaction, err := h.service.Checkout(request.Items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}

func (h *TransactionHandler) HandleTransactionsByDateRange(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetTransactionsByDateRange(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TransactionHandler) GetTransactionsByDateRange(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	var report *model.TransactionReportRequest
	var err error

	if startDate == "" || endDate == "" {
		report, err = h.service.GetTodayTransactions()
	} else {
		report, err = h.service.GetTransactionsByDateRange(startDate, endDate)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
