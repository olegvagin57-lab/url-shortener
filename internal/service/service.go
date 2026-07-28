package service

import (
	"time"
	"url-shortener/internal/generator"
	"url-shortener/internal/models"
	"url-shortener/internal/storage"
)

type Service struct {
	storage *storage.Storage
	nextID  int
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		storage: storage,
		nextID:  1,
	}
}

func (s *Service) CreateShortURL(originalURL string) (models.ShortURL, error) {
	shortURL := models.ShortURL{
		Code:        generator.GenerateCode(s.nextID),
		OriginalURL: string(originalURL),
		CreatedAt:   time.Now(),
		Clicks:      1,
	}
	s.nextID++
	err := s.storage.Add(shortURL)
	if err != nil {
		return models.ShortURL{}, err
	}
	return shortURL, nil
}

func (s *Service) RedirectToURL(code string) (models.ShortURL, error) {
	urlToRedirect, err := s.storage.Get(code)
	if err != nil {
		return models.ShortURL{}, err
	}
	return urlToRedirect, nil
}
