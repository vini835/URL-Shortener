package main

import (
	"log"
	"net/http"

	"url-shortener/internal/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/shorten", handler.ShortenURL).Methods("POST")

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
