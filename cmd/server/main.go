package main

import (
	"fmt"
	"net/http"
	"time"
	"log"
)

var storage Storage

type ShortURL struct{
	Code string
	Original string
	CreatedAt time.Time
	Clicks int
}

type Storage struct{
	URLs map[string]ShortURL
}

func (s *Storage) Add(shortURL ShortURL) error{
	if s.URLs == nil{
		return fmt.Errorf("Map is not created")
	}
	_, ok := s.URLs[shortURL.Code]
	if ok{
		err := fmt.Errorf("ShortURL already exist")
		return err
	}else{
		s.URLs[shortURL.Code] = shortURL
	}
	fmt.Println(s.URLs[shortURL.Code])
	return nil
}

func mainHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Test")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод неподдерживается", http.StatusMethodNotAllowed)
		return
	}
	name := r.URL.Query().Get("name")
	// fmt.Println(r.URL.Query())
	if name != "" {
		fmt.Fprintf(w, "Привет, %s", name)
	} else {
		fmt.Fprintln(w, "Привет, гость")
	}
	err:= storage.Add(ShortURL{
		Code: "2123",
		Original: "google.com",
		CreatedAt: time.Now(),
		Clicks: 1,
		}); if err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
}

func main() {
	http.HandleFunc("/", mainHandler)

	http.HandleFunc("/hello", helloHandler)

	storage.URLs = make(map[string]ShortURL)

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}
