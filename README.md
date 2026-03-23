# Library API

Simple REST API for managing a library catalog while practicing Go back-end fundamentals: HTTP routing with `net/http`, request handlers, database connectivity, and CRUD workflows on top of MySQL.

## What You Learn Here
- Wiring environment-driven MySQL connections (`DSN` via `.env`).
- Designing a minimal data model (`models.Book`).
- Implementing repository helpers for CRUD with `database/sql`.
- Building HTTP handlers that decode/encode JSON payloads.
- Registering method-aware routes using Go 1.22+ patterns and `http.ServeMux`.

## Project Structure
```text
cmd/api/main.go        # bootstrapper: load env, run server
internal/router        # http.ServeMux wiring
internal/handlers      # book CRUD handlers
internal/db            # DB connection + queries
internal/models        # Book struct definition
```

## Prerequisites
- Go 1.22+
- MySQL 8.x (or compatible)
- `github.com/go-sql-driver/mysql`
- `github.com/joho/godotenv`

## Setup
1. Copy `.env.example` (or create a new `.env`) and define `DSN` in the Go MySQL driver format, e.g.:
   ```env
   DSN="user:password@tcp(127.0.0.1:3306)/library?parseTime=true"
   ```
2. Create the `books` table:
   ```sql
   CREATE TABLE books (
     id INT AUTO_INCREMENT PRIMARY KEY,
     title VARCHAR(255) NOT NULL,
     author VARCHAR(255) NOT NULL,
     genre VARCHAR(100) NOT NULL,
     published_year INT NOT NULL,
     available BOOLEAN DEFAULT TRUE,
     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
   );
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```

## Run the API
```bash
go run ./cmd/api
```
Server listens on `:8080`.

## Endpoints & cURL Examples

### Create Book
```bash
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{
        "title": "The Pragmatic Programmer",
        "author": "Andrew Hunt",
        "genre": "Tech",
        "published_year": 1999,
        "available": true
      }'
```

### List Books
```bash
curl http://localhost:8080/books
```

### Get Book by ID
```bash
curl http://localhost:8080/books/1
```

### Update Book
```bash
curl -X PUT http://localhost:8080/books/1 \
  -H "Content-Type: application/json" \
  -d '{
        "title": "Pragmatic Programmer (20th)",
        "author": "Andy Hunt",
        "genre": "Tech",
        "published_year": 2019,
        "available": true
      }'
```

### Delete Book
```bash
curl -X DELETE http://localhost:8080/books/1
```

## Next Ideas
1. Add validation for request payloads (title length, year bounds).
2. Swap `database/sql` helpers for a repository interface + tests.
3. Add pagination/filtering to `GET /books`.
4. Dockerize the API and the MySQL instance for quicker setup.
