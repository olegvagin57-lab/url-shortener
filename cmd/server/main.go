package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	storage    Storage
	urlCounter int = 100
)

type ShortURL struct {
	Code        string
	OriginalURL string
	CreatedAt   time.Time
	Clicks      int
}

type Storage struct {
	URLs map[string]ShortURL
}

func (s *Storage) Add(shortURL ShortURL) error {
	if s.URLs == nil {
		return fmt.Errorf("Map is not created")
	}
	_, ok := s.URLs[shortURL.Code]
	if ok {
		return fmt.Errorf("ShortURL already exist")
	}
	s.URLs[shortURL.Code] = shortURL
	fmt.Println(s.URLs[shortURL.Code])
	return nil
}

func (s *Storage) Get(code string) (ShortURL, error) {
	url, ok := s.URLs[code]
	if ok {
		return url, nil
	} else {
		return ShortURL{}, fmt.Errorf("Short URL is not exist")
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод неподдерживается", http.StatusMethodNotAllowed)
		return
	} else {
		userUrl, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Read URL error", http.StatusInternalServerError)
			return
		}
		urlCounter++
		shortCode := fmt.Sprintf("%sURL", strconv.Itoa(urlCounter))
		err = storage.Add(ShortURL{
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
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод неподдерживается", http.StatusMethodNotAllowed)
		return
	} else {
		code := strings.TrimPrefix(r.URL.Path, "/redirect/")
		urlToRedirect, err := storage.Get(code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, urlToRedirect.OriginalURL, http.StatusSeeOther)
	}
}

func main() {
	storage.URLs = make(map[string]ShortURL)

	http.HandleFunc("/", postHandler)

	http.HandleFunc("/redirect/", getHandler)

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}
