package handler

import (
	"encoding/json"
	"net/http"

	"url-shortener/internal/service"
	"url-shortener/pkg"
)

var shortenerService = service.NewShortenerService()

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req pkg.ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	shortURL := shortenerService.Shorten(req.URL)
	json.NewEncoder(w).Encode(pkg.ShortenResponse{ShortURL: shortURL})
}
