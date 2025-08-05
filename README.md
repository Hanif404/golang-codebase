# golang-codebase

A Go web application built with Fiber framework and MySQL database.

## Features

- REST API with Fiber framework
- MySQL database integration
- Clean architecture (domain, service, repository pattern)
- Environment configuration

## Setup

1. Install dependencies:
```bash
go mod tidy
```

2. Configure environment:
```bash
cp .env.example .env
# Edit .env with your database credentials
```

3. Run the application:
```bash
go run src/cmd/app/main.go
```

The server will start on `http://localhost:3000`

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