package main

import (
	"log"
	"net/http"
	"url-shortener/internal/handlers"
	"url-shortener/internal/storage"
)

func main() {
	storage := storage.Storage{}

	storage.InitStorage()

	handler := handlers.NewHandler(&storage)

	http.HandleFunc("/", handler.PostHandler)

	http.HandleFunc("/redirect/", handler.GetHandler)

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}
