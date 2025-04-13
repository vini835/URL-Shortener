package handler

import (
	"encoding/json"
	"net/http"

	"url-shortener/internal/service"
	"url-shortener/pkg"

	"github.com/gorilla/mux"
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

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.NotFound(w, r)
		return
	}

	original := shortenerService.Resolve(id)
	if original == "" {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, original, http.StatusFound)
}

func GetTopDomains(w http.ResponseWriter, r *http.Request) {
	topDomains := shortenerService.TopDomains(3)
	json.NewEncoder(w).Encode(topDomains)
}
