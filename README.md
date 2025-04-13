# URL Shortener API (Golang)

A simple URL shortening service built in Go, featuring:

- URL shortening with deduplication
- In-memory storage
- Redirection
- Top domain metrics
- Dockerized deployment

---

## 🚀 How to Run

### ✅ Run locally (requires Go 1.20+)
```bash
go run ./cmd
```

### 🐳 Run with Docker
```bash
docker build -t url-shortener .
docker run -p 8080:8080 url-shortener
```

---

## 🔗 API Endpoints

### 1. Shorten URL
`POST /shorten`
```bash
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://Udemy.com"}'
```
Response:
```json
{"short_url": "http://localhost:8080/1"}
```

### 2. Redirect to original URL
`GET /{id}`
```bash
curl -v http://localhost:8080/1
```
Redirects to: `https://Udemy.com`

### 3. Metrics - Top 3 Shortened Domains
`GET /metrics/top`
```bash
curl http://localhost:8080/metrics/top
```
Response:
```json
{
  "example.com": 3,
  "youtube.com": 2,
  "udemy.com": 1
}
```

---

## 🧪 Tests
> (Optional — add Go test files to `internal/` folders to test services and handlers)

---

## 📦 Docker Image
- Docker Hub: [vintasharma/url-shortener](https://hub.docker.com/r/vintasharma/url-shortener)

## 🔒 GitHub
- Private repo shared with: `anju-infracloud`

---

## 🛠 Built With
- Go 1.21
- Gorilla Mux
- Standard Library

---

## 📬 Contact
If you need access to the repo or image, feel free to reach out!
