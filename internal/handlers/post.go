package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод неподдерживается", http.StatusMethodNotAllowed)
		return
	}
	req := UserOriginalURL{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	shortURL, err := h.service.CreateShortURL(req.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	codeURL := "http://localhost:8080/" + shortURL.Code
	urlResponse := CreateShortURLResponse(shortURL, codeURL)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(urlResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
