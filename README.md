# Kasir API

A simple RESTful API for managing cashier/point-of-sale operations, built with Go. This API provides endpoints for managing products and categories.

## Live Demo

The API is deployed and accessible at: **https://kasir-api-production-efe3.up.railway.app/**

Try it out:
- API Info: https://kasir-api-production-efe3.up.railway.app/
- Get all products: https://kasir-api-production-efe3.up.railway.app/api/produk
- Health check: https://kasir-api-production-efe3.up.railway.app/health

## Features

- Product management (CRUD operations)
- Category management (CRUD operations)
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
kasir-api/
├── main.go              # Main application file with HTTP handlers
├── go.mod              # Go module definition
├── internal/
│   └── model/
│       ├── produk.go   # Product model
│       └── kategori.go # Category model
└── README.md           # This file
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
    "id": 1,
    "name": "Laptop",
    "price": 15000000,
    "stock": 10, 
    "category_id": 1,
    "category": "Electronics"
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
  "id": 4,
  "name": "Smartphone",
  "price": 8000000,
  "stock": 20,
  "category_id": 1,
  "category": "Electronics"
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
  "category_id": 4
}
```

**Response:**
```json
{
  "id": 7,
  "name": "Jeans",
  "price": 400000,
  "stock": 7,
  "category_id": 4,
  "category": ""
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
  "category_id": 1,
  "category": ""
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

