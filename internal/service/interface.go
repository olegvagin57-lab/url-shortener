package service

import (
	"url-shortener/internal/models"
)

type Storage interface {
	Add(shortURL models.ShortURL) error
	Get(code string) (models.ShortURL, error)
	Update(url models.ShortURL) error
}
