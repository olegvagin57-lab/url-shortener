package handlers

import "url-shortener/internal/storage"

type Handler struct {
	Storage *storage.Storage
}
