package handlers

import "url-shortener/internal/models"

type ShortURLResponse struct {
	Code     string `json:"code"`
	ShortURL string `json:"short_url"`
}

type UserOriginalURL struct {
	URL string `json:"url"`
}

func NewShortURLResponse(shortURL models.ShortURL) ShortURLResponse {
	urlResponse := ShortURLResponse{
		Code:     shortURL.Code,
		ShortURL: "http://localhost:8080/redirect/" + shortURL.Code,
	}
	return urlResponse
}
