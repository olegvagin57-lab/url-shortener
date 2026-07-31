package handlers

import "url-shortener/internal/models"

type ShortURLResponse struct {
	Code     string `json:"code"`
	ShortURL string `json:"short_url"`
}

func CreateShortURLResponse(shortURL models.ShortURL, codeURL string) ShortURLResponse {
	urlResponce := ShortURLResponse{
		Code:     shortURL.Code,
		ShortURL: codeURL,
	}
	return urlResponce
}
