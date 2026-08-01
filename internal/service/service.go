package service

import (
	"fmt"
	"net/url"
	"strings"
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
		nextID:  100,
	}
}

func (s *Service) CreateShortURL(originalURL string) (models.ShortURL, error) {
	originalURL = strings.TrimSpace(originalURL)
	err := s.validateURL(originalURL)
	if err != nil {
		return models.ShortURL{}, err
	}
	shortURL := models.ShortURL{
		Code:        generator.GenerateCode(s.nextID),
		OriginalURL: string(originalURL),
		CreatedAt:   time.Now(),
		Clicks:      0,
	}
	s.nextID++
	err = s.storage.Add(shortURL)
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
	err = s.updateCliсks(urlToRedirect)
	if err != nil {
		return models.ShortURL{}, err
	}
	fmt.Println(urlToRedirect)
	return urlToRedirect, nil
}

func (s *Service) updateCliсks(url models.ShortURL) error {
	url.Clicks++
	err := s.storage.Update(url)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) validateURL(originalURL string) error {
	u, err := url.ParseRequestURI(originalURL)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return fmt.Errorf("Invalid URL")
	}
	return nil
}
