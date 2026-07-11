# AVA Air Task API

A simple REST API built with Go and PostgreSQL for managing tasks.

This project was developed as a backend coding assignment. It follows a layered architecture with clean separation of concerns and demonstrates REST API design, PostgreSQL integration, migrations, Docker support, and unit testing.

---

## Features

- RESTful CRUD API
- PostgreSQL database
- Layered Architecture
- Repository Pattern
- Environment Variables (.env)
- Database Migrations
- Docker Support
- Unit Tests

---

## Project Structure

```
task-api/
│
├── cmd/
│   └── api/
│
├── internal/
│   ├── config/
│   ├── database/
│   ├── handler/
│   ├── model/
│   ├── repository/
│   ├── routes/
│   └── service/
│
├── migrations/
│
├── Dockerfile
├── docker-compose.yml
├── .env.example
├── go.mod
└── README.md
```

---

## Architecture

The application follows a layered architecture.

```
HTTP Request
      │
      ▼
Handler
      │
      ▼
Service
      │
      ▼
Repository
      │
      ▼
PostgreSQL
```
---

## Environment Variables

Create a `.env` file using `.env.example`.

Example:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=task_api
DB_SSLMODE=disable
```

---

## Running Locally

Install dependencies

```bash
go mod download
```

Run the application

```bash
go run ./cmd/api
```

---

## Database Migration

Run migrations

```bash
migrate -path migrations -database "postgres://postgres:YOUR_PASSWORD@localhost:5432/task_api?sslmode=disable" up
```

Rollback

```bash
migrate -path migrations -database "postgres://postgres:YOUR_PASSWORD@localhost:5432/task_api?sslmode=disable" down
```

---

## Docker

Build and run

```bash
docker compose up --build
```

---

## Running Tests

```bash
go test ./...
```

---

## API Endpoints

| Method | Endpoint | Description |
|---------|----------|-------------|
| GET | /tasks | Get all tasks |
| GET | /tasks/{id} | Get task by id |
| POST | /tasks | Create task |
| PUT | /tasks/{id} | Update task |
| DELETE | /tasks/{id} | Delete task |

---

## Sample Request

POST /tasks

```json
{
  "title": "Learn Go",
  "description": "Practice REST API",
  "completed": false
}
```

---

## Sample Response

```json
{
  "id": 1,
  "title": "Learn Go",
  "description": "Practice REST API",
  "completed": false,
  "created_at": "2026-07-11T20:00:00Z",
  "updated_at": "2026-07-11T20:00:00Z"
}
```

---

## Technologies

- Go
- PostgreSQL
- Docker
- golang-migrate
- godotenv

---
## Requirements

- Go 1.25+
- PostgreSQL 15+
- Docker (optional)
## Author

Sepehr Daemi