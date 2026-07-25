package main

import (
	"log"
	"net/http"
	"url-shortener/internal/handlers"
	"url-shortener/internal/storage"
)

func main() {
	storage.Stor.InitStorage()

	http.HandleFunc("/", handlers.PostHandler)

	http.HandleFunc("/redirect/", handlers.GetHandler)

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}
