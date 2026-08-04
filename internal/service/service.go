package service

import (
	"fmt"
	"net/url"
	"strings"
	"time"
	"url-shortener/internal/generator"
	"url-shortener/internal/models"
)

type Service struct {
	storage   Storage
	generator *generator.Generator
}

func NewService(storage Storage, generator *generator.Generator) *Service {
	return &Service{
		storage:   storage,
		generator: generator,
	}
}

func (s *Service) CreateShortURL(originalURL string) (models.ShortURL, error) {
	originalURL = strings.TrimSpace(originalURL)
	err := s.validateURL(originalURL)
	if err != nil {
		return models.ShortURL{}, err
	}
	shortURL := models.ShortURL{
		Code:        s.generator.GenerateCode(),
		OriginalURL: string(originalURL),
		CreatedAt:   time.Now(),
		Clicks:      0,
	}
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
		return fmt.Errorf("invalid URL")
	}
	return nil
}
