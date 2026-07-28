package handlers

import (
	"net/http"
	"strings"
)

func (h *Handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод неподдерживается", http.StatusMethodNotAllowed)
		return
	}
	code := strings.TrimPrefix(r.URL.Path, "/redirect/")
	urlToRedirect, err := h.service.RedirectToURL(code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// urlToRedirect.Clicks++
	// err1 := h.storage.Update(urlToRedirect)
	// if err1 != nil {
	// 	http.Error(w, "Failed to update", http.StatusBadRequest)
	// 	return
	// }
	http.Redirect(w, r, urlToRedirect.OriginalURL, http.StatusSeeOther)
}
