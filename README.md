# Kasir API

A simple RESTful API for managing cashier/point-of-sale operations, built with Go. This API provides endpoints for managing products and categories.

## Live Demo

The API is deployed and accessible at: **https://kasir-api-production-efe3.up.railway.app/**

Try it out:
- API Info: https://kasir-api-production-efe3.up.railway.app/
- Health check: https://kasir-api-production-efe3.up.railway.app/health
- Get all products: https://kasir-api-production-efe3.up.railway.app/api/produk
- Get categories: https://kasir-api-production-efe3.up.railway.app/api/categories
- Checkout transactions: https://kasir-api-production-efe3.up.railway.app/api/checkout
- Get report transaction within given date: https://kasir-api-production-efe3.up.railway.app/api/report

## Features

- Product management (CRUD operations)
- Category management (CRUD operations)
- Checkout management
- Transaction report request
- Health check endpoint
- API information endpoint
- JSON response format
- CORS enabled for product endpoints

## Prerequisites

- Go 1.25.6 or higher

## Installation

1. Clone the repository:
```bash
git clone https://github.com/alfanugraha/kasir-api.git
cd kasir-api
```

2. Install dependencies (if any):
```bash
go mod download
```

3. Run the application:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## Deployment

This application is deployed on [Railway](https://railway.app/).

**Production URL:** https://kasir-api-production-efe3.up.railway.app/

The deployment is configured to:
- Automatically deploy from the master branch
- Run on Railway's infrastructure
- Use the default port provided by Railway via the `PORT` environment variable

## Project Structure

```
└── 📁kasir-api
    └── 📁database
        ├── config.go
    └── 📁handler
        ├── category_handler.go
        ├── product_handler.go
        ├── transaction_handler.go
    └── 📁model
        ├── category_model.go
        ├── product_model.go
        ├── transaction_model.go
    └── 📁repository
        ├── category_repository.go
        ├── product_repository.go
        ├── transaction_repository.go
    └── 📁service
        ├── category_service.go
        ├── product_service.go
        ├── transaction_service.go
    ├── .gitignore
    ├── go.mod
    ├── go.sum
    ├── main.go
    └── README.md
```

## API Endpoints

### Root

#### Get API Information
```
GET /
```
Returns API information including name, version, and available endpoints.

**Response:**
```json
{
  "name": "Kasir API",
  "version": "2.0.0",
  "endpoints": [...]
}
```

### Products

#### Get All Products
```
GET /api/produk
```
Returns a list of all products.

**Response:**
```json
[
  {
    "id": 3,
    "name": "Laptop",
    "price": 15000000,
    "stock": 3,
    "category": {
        "id": 1,
        "category": "Electronics",
        "description": "Electronic devices and gadgets"
    }
]
```

#### Get Product by ID
```
GET /api/produk/{id}
```
Returns a single product by ID.

**Response:**
```json
{
    "id": 3,
    "name": "Laptop",
    "price": 15000000,
    "stock": 3,
    "category": {
        "id": 1,
        "category": "Electronics",
        "description": "Electronic devices and gadgets"
    }
}
```

#### Create New Product
```
POST /api/produk
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Jeans",
  "price": 400000,
  "stock": 7,
  "category_id": 5
}
```

**Response:**
```json
{
  "id": 7,
  "name": "Jeans",
  "price": 400000,
  "stock": 7,
  "category": {
    "id": 5,
    "category": "Clothing",
    "description": "Apparel and fashion items"
  }
}
```

#### Update Product
```
PUT /api/produk/{id}
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Laptop",
  "price": 15000000,
  "stock": 3,
  "category_id": 1
}
```

**Response:**
```json
{
  "id": 3,
  "name": "Laptop",
  "price": 15000000,
  "stock": 3,
    "category": {
        "id": 1,
        "category": "Electronics",
        "description": "Electronic devices and gadgets"
    }
}
```

#### Delete Product
```
DELETE /api/produk/{id}
```

**Response:**
```json
{
    "message": "Product deleted successfully"
}
```

### Categories

#### Get All Categories
```
GET /api/categories
```
Returns a list of all categories.

**Response:**
```json
[
  {
    "id": 1,
    "category": "Elektronik",
    "description": "Perangkat elektronik seperti laptop, smartphone, dll."
  }
]
```

#### Get Category by ID
```
GET /api/categories/{id}
```
Returns a single category by ID.

**Response:**
```json
{
  "id": 1,
  "category": "Elektronik",
  "description": "Perangkat elektronik seperti laptop, smartphone, dll."
}
```

#### Create New Category
```
POST /api/categories
Content-Type: application/json
```

**Request Body:**
```json
{
  "category": "Makanan",
  "description": "Produk makanan dan minuman"
}
```

**Response:**
```json
{
  "id": 3,
  "category": "Makanan",
  "description": "Produk makanan dan minuman"
}
```

#### Update Category
```
PUT /api/categories/{id}
Content-Type: application/json
```

**Request Body:**
```json
{
  "category": "Elektronik Rumah",
  "description": "Perangkat elektronik untuk rumah tangga"
}
```

**Response:**
```json
{
  "id": 1,
  "category": "Elektronik Rumah",
  "description": "Perangkat elektronik untuk rumah tangga"
}
```

#### Delete Category
```
DELETE /api/categories/{id}
```

**Response:**
```json
{
  "message": "Category deleted successfully"
}
```

### Transaction

#### Checkout
```
POST /api/categories
Content-Type: application/json
```

**Request Body:**
```json
{
    "items": [
        {
            "product_id": 5,
            "quantity": 6
        },
        {
            "product_id": 7,
            "quantity": 2
        }
    ]
}
```

**Response:**
```json
{
    "id": 1,
    "total_price": 3800000,
    "created_at": "0001-01-01T00:00:00Z",
    "details": [
        {
            "id": 0,
            "transaction_id": 0,
            "product_id": 5,
            "product_name": "T-Shirt",
            "quantity": 6,
            "subtotal": 3000000
        },
        {
            "id": 0,
            "transaction_id": 0,
            "product_id": 7,
            "product_name": "Jeans",
            "quantity": 2,
            "subtotal": 800000
        }
    ]
}
```

#### Today's Transaction Report
```
GET /api/report/hari-ini
```
Returns a report of today's transaction.

**Response:**
```json
{
    "total_revenue": 3800000,
    "total_transaksi": 1,
    "produk_terlaris": {
        "nama": "T-Shirt",
        "qty_terjual": 6
    }
}
```

#### Request Transaction Report By Date Range
```
GET /api/report?start_date=2025-01-01&end_date=2026-12-31
```
Returns a report of today's transaction.

**Response:**
```json
{
    "total_revenue": 3800000,
    "total_transaksi": 1,
    "produk_terlaris": {
        "nama": "T-Shirt",
        "qty_terjual": 6
    }
}
```

### Health Check

#### Check API Health
```
GET /health
```

**Response:**
```json
{
  "status": "OK",
  "message": "API is running smoothly"
}
```

## Testing with cURL

### Using Local Server (http://localhost:8080)

### Get all products:
```bash
curl http://localhost:8080/api/produk
```

### Get product by ID:
```bash
curl http://localhost:8080/api/produk/1
```

### Create a new product:
```bash
curl -X POST http://localhost:8080/api/produk \
  -H "Content-Type: application/json" \
  -d '{"nama":"Keyboard","harga":150000,"stok":50,"id_category":1}'
```

### Update a product:
```bash
curl -X PUT http://localhost:8080/api/produk/1 \
  -H "Content-Type: application/json" \
  -d '{"nama":"Laptop Updated","harga":16000000,"stok":8,"id_category":1}'
```

### Delete a product:
```bash
curl -X DELETE http://localhost:8080/api/produk/1
```

### Using Production Server (Railway)

### Get all products:
```bash
curl https://kasir-api-production-efe3.up.railway.app/api/produk
```

### Get product by ID:
```bash
curl https://kasir-api-production-efe3.up.railway.app/api/produk/1
```

### Create a new product:
```bash
curl -X POST https://kasir-api-production-efe3.up.railway.app/api/produk \
  -H "Content-Type: application/json" \
  -d '{"nama":"Keyboard","harga":150000,"stok":50,"id_category":1}'
```

### Update a product:
```bash
curl -X PUT https://kasir-api-production-efe3.up.railway.app/api/produk/1 \
  -H "Content-Type: application/json" \
  -d '{"nama":"Laptop Updated","harga":16000000,"stok":8,"id_category":1}'
```

### Delete a product:
```bash
curl -X DELETE https://kasir-api-production-efe3.up.railway.app/api/produk/1
```

## Data Models

### Produk (Product)
```go
type Produk struct {
    ID          int     `json:"id"`
    Nama        string  `json:"nama"`
    Harga       float64 `json:"harga"`
    Stok        int     `json:"stok"`
}
```

### Kategori (Category)
```go
type Kategori struct {
    ID         int    `json:"id"`
    Kategori   string `json:"nama"`
    Keterangan string `json:"description"`
}
```

## Error Responses

The API returns standard HTTP status codes:

- `200 OK` - Request successful
- `201 Created` - Resource created successfully
- `400 Bad Request` - Invalid request data
- `404 Not Found` - Resource not found
- `500 Internal Server Error` - Server error

Error response format:
```
HTTP/1.1 404 Not Found
Content-Type: text/plain

Produk belum ada
```

## Development Notes

- Data is stored in memory (not persistent)
- Server restarts will reset all data to initial state
- CORS is enabled for product endpoints
- Auto-incremental IDs for new resources

## License

This project is created for educational purposes.

