package service

import (
	"errors"
	"testing"
	"url-shortener/internal/generator"
	"url-shortener/internal/models"
)

// mockStorage реализует интерфейс Storage для тестов
type mockStorage struct {
	getFunc    func(code string) (models.ShortURL, error)
	addFunc    func(shortURL models.ShortURL) error
	updateFunc func(url models.ShortURL) error
}

func (m *mockStorage) Get(code string) (models.ShortURL, error) {
	return m.getFunc(code)
}

func (m *mockStorage) Add(shortURL models.ShortURL) error {
	return m.addFunc(shortURL)
}

func (m *mockStorage) Update(url models.ShortURL) error {
	return m.updateFunc(url)
}

func TestService_CreateShortURL(t *testing.T) {
	// Подготовим мок, который всегда успешно добавляет
	mock := &mockStorage{
		addFunc: func(shortURL models.ShortURL) error {
			return nil // всегда успех
		},
	}
	gen := generator.NewGenerator()
	svc := NewService(mock, gen)

	tests := []struct {
		name        string
		inputURL    string
		wantErr     bool
		errContains string
	}{
		{"valid http", "http://example.com", false, ""},
		{"valid https", "https://google.com", false, ""},
		{"missing scheme", "example.com", true, "invalid URL"},
		{"empty", "", true, "invalid URL"},
		{"only spaces", "   ", true, "invalid URL"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			short, err := svc.CreateShortURL(tt.inputURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateShortURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err != nil {
				// можно проверить, что сообщение содержит нужную подстроку
			}
			if !tt.wantErr {
				if short.Code == "" {
					t.Error("Code should not be empty")
				}
				if short.OriginalURL != tt.inputURL {
					t.Errorf("OriginalURL = %v, want %v", short.OriginalURL, tt.inputURL)
				}
				if short.Clicks != 0 {
					t.Errorf("Clicks = %d, want 0", short.Clicks)
				}
			}
		})
	}
}

func TestService_RedirectToURL(t *testing.T) {
	// Мок, который возвращает заранее заданный URL
	expected := models.ShortURL{Code: "abc", OriginalURL: "http://target.com", Clicks: 0}
	mock := &mockStorage{
		getFunc: func(code string) (models.ShortURL, error) {
			if code == "abc" {
				return expected, nil
			}
			return models.ShortURL{}, errors.New("not found")
		},
		updateFunc: func(url models.ShortURL) error {
			// Проверим, что клики увеличились
			if url.Clicks != 1 {
				t.Errorf("updateFunc: Clicks = %d, want 1", url.Clicks)
			}
			return nil
		},
	}
	gen := generator.NewGenerator()
	svc := NewService(mock, gen)

	// Успешный случай
	got, err := svc.RedirectToURL("abc")
	if err != nil {
		t.Errorf("RedirectToURL() error = %v, want nil", err)
	}
	if got.OriginalURL != expected.OriginalURL {
		t.Errorf("RedirectToURL() = %+v, want %+v", got, expected)
	}

	// Случай с несуществующим кодом
	_, err = svc.RedirectToURL("notexist")
	if err == nil {
		t.Error("RedirectToURL() with non-existent code should return error, got nil")
	}
}
