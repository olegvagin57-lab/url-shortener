package storage

import (
	"testing"
	"url-shortener/internal/models"
)

func TestMemoryStorage_Add(t *testing.T) {
	s := &MemoryStorage{}
	s.InitStorage()

	url := models.ShortURL{Code: "abc", OriginalURL: "http://example.com"}

	// Первое добавление - успех
	err := s.Add(url)
	if err != nil {
		t.Errorf("Add() error = %v, want nil", err)
	}

	// Повторное добавление с тем же кодом - ошибка
	err = s.Add(url)
	if err == nil {
		t.Error("Add() duplicate should return error, got nil")
	}
}

func TestMemoryStorage_Get(t *testing.T) {
	s := &MemoryStorage{}
	s.InitStorage()

	url := models.ShortURL{Code: "xyz", OriginalURL: "http://test.com"}
	s.Add(url)

	// Получение существующего
	got, err := s.Get("xyz")
	if err != nil {
		t.Errorf("Get() error = %v, want nil", err)
	}
	if got.Code != url.Code || got.OriginalURL != url.OriginalURL {
		t.Errorf("Get() = %+v, want %+v", got, url)
	}

	// Получение несуществующего
	_, err = s.Get("none")
	if err == nil {
		t.Error("Get() non-existent should return error, got nil")
	}
}

func TestMemoryStorage_Update(t *testing.T) {
	s := &MemoryStorage{}
	s.InitStorage()

	url := models.ShortURL{Code: "upd", OriginalURL: "http://old.com", Clicks: 5}
	s.Add(url)

	// Обновляем
	url.OriginalURL = "http://new.com"
	url.Clicks = 10
	err := s.Update(url)
	if err != nil {
		t.Errorf("Update() error = %v, want nil", err)
	}

	// Проверяем, что обновилось
	got, _ := s.Get("upd")
	if got.OriginalURL != "http://new.com" || got.Clicks != 10 {
		t.Errorf("After Update() got = %+v, want OriginalURL=new.com, Clicks=10", got)
	}
}
