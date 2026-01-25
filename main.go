package main

import (
	"encoding/json"
	"fmt"
	"kasir-api/internal/model"
	"net/http"
	"strconv"
	"strings"
)

var produk = []model.Produk{
	{ID: 1, Nama: "Laptop", Harga: 15000000, Stok: 10},
	{ID: 2, Nama: "Smartphone", Harga: 5000000, Stok: 25},
	{ID: 3, Nama: "Tablet", Harga: 7000000, Stok: 15},
}

func getProdukByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	for _, p := range produk {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

func updateProdukByID(w http.ResponseWriter, r *http.Request) {
	// get id from request URL
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")

	// change int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	// get data from request body
	var updateProduk model.Produk
	err = json.NewDecoder(r.Body).Decode(&updateProduk)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	// loop produk, find id, change data
	for i := range produk {
		if produk[i].ID == id {
			updateProduk.ID = id
			produk[i] = updateProduk

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateProduk)
			return
		}
	}
}

func deleteProdukByID(w http.ResponseWriter, r *http.Request) {
	// get id from request URL
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")

	// change int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	// loop produk, find id
	for i, p := range produk {
		if p.ID == id {
			// create new slice before and after index
			produk = append(produk[:i], produk[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Produk berhasil dihapus",
			})
			return
		}
	}

	http.Error(w, "Produk belum ada", http.StatusNotFound)

}

func handleAPIInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	apiInfo := map[string]interface{}{
		"name":    "Kasir API",
		"version": "1.0.0",
		"endpoints": []map[string]string{
			{"method": "GET", "path": "/api/produk", "description": "Get all products"},
			{"method": "POST", "path": "/api/produk", "description": "Create new product"},
			{"method": "GET", "path": "/api/produk/{id}", "description": "Get product by ID"},
			{"method": "PUT", "path": "/api/produk/{id}", "description": "Update product by ID"},
			{"method": "DELETE", "path": "/api/produk/{id}", "description": "Delete product by ID"},
			{"method": "GET", "path": "/health", "description": "Health check"},
		},
	}
	json.NewEncoder(w).Encode(apiInfo)
}

func main() {
	// GET localhost:8080/
	http.HandleFunc("/", handleAPIInfo)

	// GET localhost:8080/api/produk/{id}
	// PUT localhost:8080/api/produk/{id}
	// DELETE localhost:8080/api/produk/{id}
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getProdukByID(w, r)
		} else if r.Method == "PUT" {
			updateProdukByID(w, r)
		} else if r.Method == "DELETE" {
			deleteProdukByID(w, r)
		}

	})

	// GET localhost:8080/api/produk
	// POST localhost:8080/api/produk
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produk)
		} else if r.Method == "POST" {
			// read data from request
			var produkBaru model.Produk
			err := json.NewDecoder(r.Body).Decode(&produkBaru)
			if err != nil {
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}

			// input data to produk slice
			produkBaru.ID = len(produk) + 1
			produk = append(produk, produkBaru)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated) // 201
			json.NewEncoder(w).Encode(produkBaru)
		}
	})

	// localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API is running smoothly",
		})
		w.Write([]byte("OK"))
	})
	fmt.Println("Starting server at :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed to start")
	}
}
