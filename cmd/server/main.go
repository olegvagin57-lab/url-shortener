package main

import (
	"log"
	"net/http"
	"url-shortener/internal/handlers"
	"url-shortener/internal/service"
	"url-shortener/internal/storage"
)

func main() {
	storage := storage.MemoryStorage{}

	storage.InitStorage()

	svc := service.NewService(&storage)

	handler := handlers.NewHandler(svc)

	http.HandleFunc("/", handler.PostHandler)

	http.HandleFunc("/redirect/", handler.GetHandler)

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}
