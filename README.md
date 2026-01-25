# Kasir API

A simple RESTful API for managing cashier/point-of-sale operations, built with Go. This API provides endpoints for managing products and categories.

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
  "version": "1.0.0",
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
    "nama": "Laptop",
    "harga": 15000000,
    "stok": 10
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
  "id": 1,
  "nama": "Laptop",
  "harga": 15000000,
  "stok": 10
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
  "nama": "Mouse",
  "harga": 50000,
  "stok": 100
}
```

**Response:**
```json
{
  "id": 4,
  "nama": "Mouse",
  "harga": 50000,
  "stok": 100
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
  "nama": "Laptop Gaming",
  "harga": 20000000,
  "stok": 5
}
```

**Response:**
```json
{
  "id": 1,
  "nama": "Laptop Gaming",
  "harga": 20000000,
  "stok": 5
}
```

#### Delete Product
```
DELETE /api/produk/{id}
```

**Response:**
```json
{
  "message": "Produk berhasil dihapus"
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
    "nama": "Elektronik",
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
  "nama": "Elektronik",
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
  "nama": "Makanan",
  "description": "Produk makanan dan minuman"
}
```

**Response:**
```json
{
  "id": 3,
  "nama": "Makanan",
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
  "nama": "Elektronik Rumah",
  "description": "Perangkat elektronik untuk rumah tangga"
}
```

**Response:**
```json
{
  "id": 1,
  "nama": "Elektronik Rumah",
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
  "message": "Kategori berhasil dihapus"
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

