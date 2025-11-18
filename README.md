# go-url-shortener

A TinyURL-style URL shortening service built in Go with redirecting, analytics, and database storage.

## Tech Stack

- Go 1.21+
- Gin (HTTP router)
- GORM (ORM)
- SQLite (default)

## Project Structure

```
go-url-shortener/
├── cmd/
│   └── server/
│       └── main.go          # Entry point
├── models/
│   └── url.go               # URL database model
├── handlers/
│   └── handlers.go          # HTTP handlers (POST /shorten, GET /:code)
├── database/
│   └── database.go          # DB connection & initialization
├── utils/
│   └── code.go              # Code generation utility
└── go.mod                   # Go module file
```

## Setup & Installation

### Prerequisites

1. Install Go from [golang.org](https://golang.org/dl/)
2. Verify installation: `go version`

### Steps

1. **Install dependencies:**
   ```bash
   go mod download
   ```

2. **Build the server:**
   ```bash
   go build ./cmd/server
   ./server
   ```

   Or run directly:
   ```bash
   go run ./cmd/server/main.go
   ```

3. **Test the API:**

   Create a short URL:
   ```bash
   curl -X POST http://localhost:8080/shorten \
     -H "Content-Type: application/json" \
     -d '{"url": "https://example.com/very/long/url"}'
   ```

   Access the short URL:
   ```bash
   curl http://localhost:8080/abc123
   # or open in browser - it will redirect
   ```

## Implementation Steps (Learning Guide)

### Step 1: Database Setup ✓
- [x] Initialize Go module
- [ ] Complete `database/database.go` - connect to SQLite and run migrations

### Step 2: Code Generation
- [ ] Complete `utils/code.go` - implement `GenerateCode()` function

### Step 3: URL Model
- [x] Created `models/url.go` with URL struct
- [ ] Verify GORM tags are correct

### Step 4: Handlers
- [ ] Complete `handlers/handlers.go`:
  - [ ] Implement `Shorten()` - create short URLs
  - [ ] Implement `Redirect()` - redirect and increment visits

### Step 5: Main Server
- [ ] Complete `cmd/server/main.go`:
  - [ ] Initialize database
  - [ ] Register routes
  - [ ] Start server

### Step 6: Testing
- [ ] Test with curl/Postman
- [ ] Verify redirects work
- [ ] Verify visit counts increment

## Features (MVP)

- ✅ Project structure setup
- ⏳ `POST /shorten` — create a short code for a long URL
- ⏳ `GET /:code` — redirect to original URL and increment visits
- ⏳ Persistent storage of mappings and visit counts

## Bonus Features (Future)

- Rate limiting
- Custom aliases
- Expiration dates
- Admin dashboard
- Dockerfile & compose for easy deployment

## Learning Resources

- [Go Tour](https://go.dev/tour/)
- [Gin Framework Docs](https://gin-gonic.com/docs/)
- [GORM Docs](https://gorm.io/docs/)
