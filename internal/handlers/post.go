package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func (h *Handler) PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод неподдерживается", http.StatusMethodNotAllowed)
		return
	}
	userUrl, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Read URL error", http.StatusInternalServerError)
		return
	}
	shortURL, err := h.service.CreateShortURL(string(userUrl))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, shortURL)
}
