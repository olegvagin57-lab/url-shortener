package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
	"url-shortener/internal/models"
	"url-shortener/internal/storage"
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
	storage.URLCounter++
	shortCode := fmt.Sprintf("%sURL", strconv.Itoa(storage.URLCounter))
	err = h.storage.Add(models.ShortURL{
		Code:        shortCode,
		OriginalURL: string(userUrl),
		CreatedAt:   time.Now(),
		Clicks:      0,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
