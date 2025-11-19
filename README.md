# go-url-shortener

A TinyURL-style URL shortening service built in Go with redirecting, analytics, and database storage.

## Tech Stack

- Go 1.21+
- Gin (HTTP router)
- GORM (ORM)
- SQLite (pure Go driver - no CGO required)

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
├── web/
│   └── index.html           # Web frontend UI
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

2. **Run the server:**
   ```bash
   go run ./cmd/server/main.go
   ```

   Or build and run:
   ```bash
   go build ./cmd/server
   ./server
   ```

3. **Access the web interface:**
   
   Open your browser and navigate to:
   ```
   http://localhost:8080
   ```
   
   You'll see a beautiful web interface where you can:
   - Enter a long URL
   - Click "Shorten URL" to create a short link
   - Copy the short URL with one click
   - Test redirects by visiting the short URL

4. **Test the API (alternative to web UI):**

   Create a short URL:
   ```bash
   # Using curl
   curl -X POST http://localhost:8080/shorten \
     -H "Content-Type: application/json" \
     -d '{"url": "https://example.com/very/long/url"}'
   
   # Using PowerShell
   Invoke-RestMethod -Uri http://localhost:8080/shorten -Method Post -ContentType "application/json" -Body '{"url":"https://example.com"}'
   ```

   Response:
   ```json
   {
     "code": "Mu4DZ6",
     "short_url": "http://localhost:8080/Mu4DZ6"
   }
   ```

   Access the short URL:
   ```bash
   # Open in browser or use curl
   curl http://localhost:8080/Mu4DZ6
   # Browser will automatically redirect to the original URL
   ```

## Implementation Status

### ✅ Completed
- [x] Database setup with SQLite (pure Go driver)
- [x] Code generation utility
- [x] URL model with GORM
- [x] Shorten handler (POST /shorten)
- [x] Redirect handler (GET /:code) with visit tracking
- [x] Main server setup and routing
- [x] Web frontend UI (HTML/CSS/JavaScript)
- [x] Testing and verification

## Features

- ✅ **Web Interface** — Beautiful, modern UI for shortening URLs
- ✅ `POST /shorten` — create a short code for a long URL
- ✅ `GET /:code` — redirect to original URL and increment visits
- ✅ Persistent storage of mappings and visit counts
- ✅ Automatic collision handling for short codes
- ✅ Visit counter tracking
- ✅ Pure Go implementation (no CGO required)
- ✅ No npm/build tools required — simple HTML/CSS/JS frontend

## API Endpoints

### POST /shorten
Creates a short URL from a long URL.

**Request:**
```json
{
  "url": "https://example.com/very/long/url"
}
```

**Response:**
```json
{
  "code": "abc123",
  "short_url": "http://localhost:8080/abc123"
}
```

### GET /:code
Redirects to the original URL and increments the visit counter.

**Response:** HTTP 302 Redirect to the original URL

### GET /
Serves the web interface for easy URL shortening.

**Response:** HTML page with interactive UI

## Future Enhancements

- Rate limiting
- Custom aliases
- Expiration dates
- Admin dashboard
- Analytics endpoint
- Dockerfile & compose for easy deployment

## Learning Resources

- [Go Tour](https://go.dev/tour/)
- [Gin Framework Docs](https://gin-gonic.com/docs/)
- [GORM Docs](https://gorm.io/docs/)
