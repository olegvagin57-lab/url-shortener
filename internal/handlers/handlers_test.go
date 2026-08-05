package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"url-shortener/internal/models"
)

// mockService реализует ServiceInterface
type mockService struct {
	createShortURLFunc func(originalURL string) (models.ShortURL, error)
	redirectToURLFunc  func(code string) (models.ShortURL, error)
}

func (m *mockService) CreateShortURL(originalURL string) (models.ShortURL, error) {
	return m.createShortURLFunc(originalURL)
}

func (m *mockService) RedirectToURL(code string) (models.ShortURL, error) {
	return m.redirectToURLFunc(code)
}

func TestHandler_PostHandler(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    string
		mockCreateFunc func(originalURL string) (models.ShortURL, error)
		wantStatus     int
		wantBodyFields map[string]string // проверяем наличие полей в JSON
	}{
		{
			name:        "success",
			requestBody: `{"url":"http://example.com"}`,
			mockCreateFunc: func(originalURL string) (models.ShortURL, error) {
				return models.ShortURL{Code: "abc123", OriginalURL: originalURL}, nil
			},
			wantStatus: http.StatusCreated,
			wantBodyFields: map[string]string{
				"code":      "abc123",
				"short_url": "http://localhost:8080/redirect/abc123",
			},
		},
		{
			name:           "invalid json",
			requestBody:    `{"url":`,
			mockCreateFunc: nil, // не будет вызван
			wantStatus:     http.StatusBadRequest,
		},
		{
			name:        "service error",
			requestBody: `{"url":"http://example.com"}`,
			mockCreateFunc: func(originalURL string) (models.ShortURL, error) {
				return models.ShortURL{}, errors.New("invalid url")
			},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &mockService{
				createShortURLFunc: tt.mockCreateFunc,
			}
			handler := NewHandler(svc)

			req := httptest.NewRequest("POST", "/", strings.NewReader(tt.requestBody))
			w := httptest.NewRecorder()

			handler.PostHandler(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			if resp.StatusCode != tt.wantStatus {
				t.Errorf("status = %v, want %v", resp.StatusCode, tt.wantStatus)
			}

			if tt.wantStatus == http.StatusCreated {
				var jsonResp map[string]interface{}
				if err := json.NewDecoder(resp.Body).Decode(&jsonResp); err != nil {
					t.Fatalf("cannot decode response: %v", err)
				}
				for key, wantVal := range tt.wantBodyFields {
					gotVal, ok := jsonResp[key]
					if !ok {
						t.Errorf("missing field %q in response", key)
					} else if gotVal != wantVal {
						t.Errorf("field %q = %v, want %v", key, gotVal, wantVal)
					}
				}
			}
		})
	}
}

func TestHandler_GetHandler(t *testing.T) {
	tests := []struct {
		name             string
		requestPath      string
		mockRedirectFunc func(code string) (models.ShortURL, error)
		wantStatus       int
		wantLocation     string // для проверки редиректа
	}{
		{
			name:        "success redirect",
			requestPath: "/redirect/abc",
			mockRedirectFunc: func(code string) (models.ShortURL, error) {
				if code == "abc" {
					return models.ShortURL{OriginalURL: "http://destination.com"}, nil
				}
				return models.ShortURL{}, errors.New("not found")
			},
			wantStatus:   http.StatusSeeOther,
			wantLocation: "http://destination.com",
		},
		{
			name:        "code not found",
			requestPath: "/redirect/unknown",
			mockRedirectFunc: func(code string) (models.ShortURL, error) {
				return models.ShortURL{}, errors.New("not found")
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name:        "empty code",
			requestPath: "/redirect/",
			mockRedirectFunc: func(code string) (models.ShortURL, error) {
				// код будет пустой строкой
				if code == "" {
					return models.ShortURL{}, errors.New("empty code")
				}
				return models.ShortURL{}, nil
			},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &mockService{
				redirectToURLFunc: tt.mockRedirectFunc,
			}
			handler := NewHandler(svc)

			req := httptest.NewRequest("GET", tt.requestPath, nil)
			w := httptest.NewRecorder()

			handler.GetHandler(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			if resp.StatusCode != tt.wantStatus {
				t.Errorf("status = %v, want %v", resp.StatusCode, tt.wantStatus)
			}

			if tt.wantStatus == http.StatusSeeOther {
				location := resp.Header.Get("Location")
				if location != tt.wantLocation {
					t.Errorf("Location header = %v, want %v", location, tt.wantLocation)
				}
			}
		})
	}
}
