package main

import (
	"fmt"
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Главная страница")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "About page")
}

func contactsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Contacts")
}

func abcHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод неподдерживается", http.StatusMethodNotAllowed)
		return
	}
	name := r.URL.Query().Get("name")
	fmt.Println(r.URL.Query())
	if name != "" {
		fmt.Fprintf(w, "Привет, %s", name)
	} else {
		fmt.Fprintln(w, "Привет, гость")
	}
}

func main() {
	http.HandleFunc("/", mainHandler)

	http.HandleFunc("/about", aboutHandler)

	http.HandleFunc("/contacts", contactsHandler)

	http.HandleFunc("/abc", abcHandler)

	http.HandleFunc("/hello", helloHandler)

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		fmt.Println("Ошибка запуска сервера")
	}
}
