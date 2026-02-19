package main

import (
	"encoding/json"
	"fmt"
	"kasir-api/database"
	"kasir-api/handler"
	"kasir-api/middleware"
	"kasir-api/repository"
	"kasir-api/service"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
	APIKey string `mapstructure:"APIKEY"`
}

func handleAPIInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	apiInfo := map[string]interface{}{
		"name":    "Kasir API",
		"version": "3.0.0",
		"endpoints": []map[string]string{
			{"method": "GET", "path": "/health", "description": "Health check"},
			{"method": "GET", "path": "/api/produk", "description": "Get all products"},
			{"method": "POST", "path": "/api/produk", "description": "Create new product"},
			{"method": "GET", "path": "/api/produk/{id}", "description": "Get product by ID"},
			{"method": "PUT", "path": "/api/produk/{id}", "description": "Update product by ID"},
			{"method": "DELETE", "path": "/api/produk/{id}", "description": "Delete product by ID"},
			{"method": "GET", "path": "/api/categories", "description": "Get all categories"},
			{"method": "POST", "path": "/api/categories", "description": "Create new category"},
			{"method": "GET", "path": "/api/categories/{id}", "description": "Get category by ID"},
			{"method": "PUT", "path": "/api/categories/{id}", "description": "Update category by ID"},
			{"method": "DELETE", "path": "/api/categories/{id}", "description": "Delete category by ID"},
			{"method": "POST", "path": "/api/checkout", "description": "Checkout transaction"},
			{"method": "GET", "path": "/api/report/hari-ini", "description": "Get today's transactions report"},
			{"method": "GET", "path": "/api/report?start_date={start_date}&end_date={end_date}", "description": "Get transactions report by date range"},
		},
	}
	json.NewEncoder(w).Encode(apiInfo)
}

func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		viper.SetConfigType("env")
		if err := viper.ReadInConfig(); err != nil {
			log.Printf("Error reading config file: %v\n", err)
		}
	}

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
		APIKey: viper.GetString("APIKEY"),
	}

	// Debug: Print loaded config
	// fmt.Printf("=== Config Loaded ===\n")
	// fmt.Printf("Port: '%s'\n", config.Port)
	// fmt.Printf("API Key: '%s' (length: %d)\n", config.APIKey, len(config.APIKey))
	// fmt.Printf("====================\n")

	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	apiKeyMiddleware := middleware.APIKey(config.APIKey)

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	transactionRepo := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepo)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	// setup routes
	http.HandleFunc("/", handleAPIInfo)
	http.HandleFunc("/api/produk", middleware.CORS(middleware.Logger(productHandler.HandleProducts)))
	http.HandleFunc("/api/produk/", middleware.CORS(middleware.Logger(apiKeyMiddleware(productHandler.HandleProductByID))))

	http.HandleFunc("/api/categories", middleware.CORS(middleware.Logger(categoryHandler.HandleCategories)))
	http.HandleFunc("/api/categories/", middleware.CORS(middleware.Logger(apiKeyMiddleware(categoryHandler.HandleCategoryByID))))

	http.HandleFunc("/api/checkout", middleware.CORS(middleware.Logger(apiKeyMiddleware(transactionHandler.HandleCheckout))))
	http.HandleFunc("/api/report/hari-ini", transactionHandler.HandleTransactionsByDateRange)
	http.HandleFunc("/api/report", transactionHandler.HandleTransactionsByDateRange)

	// localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API is running smoothly",
		})
		w.Write([]byte("OK"))
	})
	fmt.Println("Starting server at localhost:" + config.Port)

	err = http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		fmt.Println("Server failed to start")
	}
}
