package storage

import (
	"fmt"
	"url-shortener/internal/models"
)

type MemoryStorage struct {
	URLs map[string]models.ShortURL
}

func (s *MemoryStorage) InitStorage() {
	s.URLs = make(map[string]models.ShortURL)
}

func (s *MemoryStorage) Add(shortURL models.ShortURL) error {
	if s.URLs == nil {
		return fmt.Errorf("Map is not created")
	}
	_, ok := s.URLs[shortURL.Code]
	if ok {
		return fmt.Errorf("ShortURL already exist")
	}
	s.URLs[shortURL.Code] = shortURL
	return nil
}

func (s *MemoryStorage) Get(code string) (models.ShortURL, error) {
	url, ok := s.URLs[code]
	if ok {
		return url, nil
	} else {
		return models.ShortURL{}, fmt.Errorf("Short URL is not exist")
	}
}

func (s *MemoryStorage) Update(url models.ShortURL) error {
	if s.URLs == nil {
		return fmt.Errorf("Map is not created")
	}
	code := url.Code
	s.URLs[code] = url
	return nil
}
