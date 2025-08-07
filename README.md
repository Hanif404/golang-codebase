# golang-codebase

A Go web application built with Fiber framework and MySQL database.

## Features

- REST API with Fiber framework
- MySQL database integration
- Clean architecture (domain, service, repository pattern)
- User management with CRUD operations
- JWT authentication
- Request validation
- Environment configuration

## Setup

1. Install dependencies:
```bash
go mod tidy
```

2. Configure environment:
```bash
# Create .env file with database credentials
DB_USER=your_username
DB_PASS=your_password
DB_HOST=localhost:3306
DB_NAME=your_database
```

3. Run the application:
```bash
go run src/cmd/app/main.go
```

The server will start on `http://localhost:3000`

## API Endpoints

### Authentication
- `POST /api/v1/login` - User login
- `POST /api/v1/registration` - User registration

### Users (Protected - Requires JWT Token)
- `GET /api/v1/users` - Get all users
- `GET /api/v1/user/:id` - Get user by ID
- `POST /api/v1/user` - Create new user
- `PUT /api/v1/user/:id` - Update user
- `DELETE /api/v1/user/:id` - Delete user

### Authentication Schema
```json
// Login
{
  "email": "john@example.com",
  "password": "password123"
}

// Registration
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

### User Schema
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

### Validation Rules
- `name`: Required, minimum 3 characters
- `email`: Required, valid email format
- `password`: Required

### Authentication
For protected endpoints, include JWT token in Authorization header:
```
Authorization: Bearer <your_jwt_token>
```

### Response Format
- Success: Returns user object or array
- Error: `{"error": "error message"}`
- Validation Error: `{"validation_errors": {"field": "rule"}}`

## Project Structure

```
src/
├── cmd/app/          # Application entry point
├── configs/          # Configuration management
├── domain/           # Domain entities
├── internal/         # Internal application logic
├── pkg/              # Shared packages
└── utils/            # Utility functions
```