# Todo Backend API 📝

Một REST API đơn giản để quản lý danh sách công việc (Todo List) được xây dựng bằng Go với framework Gin và PostgreSQL database.

## 🚀 Tính năng

- ✅ Tạo, đọc, cập nhật, xóa (CRUD) các todo items
- 📊 Lấy danh sách tất cả todos với thống kê số lượng
- 🔍 Lấy thông tin chi tiết một todo theo ID
- ✏️ Cập nhật trạng thái hoàn thành và thông tin todo
- 🗑️ Xóa todo items
- 🌐 Hỗ trợ CORS cho frontend integration
- 💾 Sử dụng PostgreSQL làm database chính
- 🔧 Auto-migration với GORM

## 🛠️ Công nghệ sử dụng

- **Backend Framework**: [Gin](https://github.com/gin-gonic/gin) - High-performance HTTP web framework
- **Database**: PostgreSQL
- **ORM**: [GORM](https://gorm.io/) - The fantastic ORM library for Golang
- **Environment Variables**: [godotenv](https://github.com/joho/godotenv)
- **CORS**: [gin-contrib/cors](https://github.com/gin-contrib/cors)

## 📁 Cấu trúc dự án

```
todo-backend/
├── controllers/           # Xử lý logic nghiệp vụ
│   └── todo_controller.go
├── database/             # Kết nối và cấu hình database
│   └── connection.go
├── middleware/           # Middleware (CORS, etc.)
│   └── cors.go
├── models/              # Định nghĩa data models
│   └── todo.go
├── routes/              # Định nghĩa API routes
│   └── routes.go
├── main.go              # Entry point của ứng dụng
├── go.mod              # Go module dependencies
├── go.sum              # Dependency checksums
└── .env                # Environment variables (cần tạo)
```

## ⚡ Cài đặt và chạy

### Yêu cầu hệ thống

- Go 1.25.1 hoặc mới hơn
- PostgreSQL database
- Git

### 1. Clone repository

```bash
git clone <your-repository-url>
cd todo-backend
```

### 2. Cài đặt dependencies

```bash
go mod tidy
```

### 3. Cấu hình environment variables

Tạo file `.env` trong thư mục root với nội dung:

```env
# Database Configuration
DB_HOST=localhost
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=todo_db
DB_PORT=5432
DB_SSLMODE=disable

# Server Configuration
PORT=8080
```

### 4. Tạo database

Tạo database PostgreSQL:

```sql
CREATE DATABASE todo_db;
```

### 5. Chạy ứng dụng

```bash
go run main.go
```

Server sẽ chạy tại `http://localhost:8080`

## 📚 API Documentation

### Health Check

```http
GET /
```

**Response:**
```json
{
  "message": "🚀 Todo API is running!",
  "version": "1.0.0",
  "database": "PostgreSQL",
  "status": "healthy"
}
```

### Todo Endpoints

Base URL: `http://localhost:8080/api`

#### 1. Lấy tất cả todos

```http
GET /api/todos
```

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "title": "Học Go programming",
      "description": "Hoàn thành tutorial Go cơ bản",
      "completed": false,
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-01T10:00:00Z"
    }
  ],
  "count": 1
}
```

#### 2. Lấy todo theo ID

```http
GET /api/todos/:id
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "title": "Học Go programming",
    "description": "Hoàn thành tutorial Go cơ bản",
    "completed": false,
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

#### 3. Tạo todo mới

```http
POST /api/todos
```

**Request Body:**
```json
{
  "title": "Học Go programming",
  "description": "Hoàn thành tutorial Go cơ bản",
  "completed": false
}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "title": "Học Go programming",
    "description": "Hoàn thành tutorial Go cơ bản",
    "completed": false,
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

#### 4. Cập nhật todo

```http
PUT /api/todos/:id
```

**Request Body:**
```json
{
  "title": "Học Go programming - Updated",
  "description": "Hoàn thành tutorial Go nâng cao",
  "completed": true
}
```

#### 5. Xóa todo

```http
DELETE /api/todos/:id
```

**Request Body:**
```json
{
  "id": 1
}
```

## 🔧 Development

### Chạy trong môi trường development

```bash
# Với live reload (cần cài gin)
go install github.com/cosmtrek/air@latest
air
```

### Build production

```bash
go build -o todo-backend main.go
./todo-backend
```

## 📝 Model Schema

### Todo Model

```go
type Todo struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    Title       string    `json:"title" gorm:"not null"`
    Description string    `json:"description"`
    Completed   bool      `json:"completed" gorm:"default:false"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

## 🛡️ Error Handling

API trả về các status codes sau:

- `200 OK` - Request thành công
- `201 Created` - Tạo resource thành công
- `400 Bad Request` - Request không hợp lệ
- `404 Not Found` - Resource không tìm thấy
- `500 Internal Server Error` - Lỗi server

**Error Response Format:**
```json
{
  "success": false,
  "message": "Error message",
  "error": "Detailed error description"
}
```

