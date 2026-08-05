package handlers

import (
	"url-shortener/internal/models"
)

type ServiceInterface interface {
	CreateShortURL(originalURL string) (models.ShortURL, error)
	RedirectToURL(code string) (models.ShortURL, error)
}
