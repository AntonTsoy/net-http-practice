package main

import (
	"fmt"
	"net/http"
)

func languageHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("lang")
	if err != nil {
		cookie = &http.Cookie{Name: "lang", Value: "en"}
		http.SetCookie(w, cookie)
	}

	switch cookie.Value {
	case "ru":
		fmt.Fprintf(w, "Привет!\n")
	default:
		fmt.Fprintf(w, "Hello!\n")
	}
}

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(languageHandler))
}
