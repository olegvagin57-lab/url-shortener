package handlers

import (
	"net/http"
	"strings"
	"url-shortener/internal/storage"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод неподдерживается", http.StatusMethodNotAllowed)
		return
	}
	code := strings.TrimPrefix(r.URL.Path, "/redirect/")
	urlToRedirect, err := storage.Stor.Get(code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	urlToRedirect.Clicks++
	err1 := storage.Stor.Update(urlToRedirect)
	if err1 != nil {
		http.Error(w, "Failed to update", http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, urlToRedirect.OriginalURL, http.StatusSeeOther)
}
