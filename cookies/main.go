package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func generateId() string {
	symbols := make([]byte, 10)
	for i := range symbols {
		switch rand.Intn(3) {
		case 0:
			symbols[i] = byte(rand.Intn(int('z' - 'a') + 1)) + 'a'
		case 1:
			symbols[i] = byte(rand.Intn(int('Z' - 'A') + 1)) + 'A'
		case 2:
			symbols[i] = byte(rand.Intn(int('9' - '0') + 1)) + '0'
		}
	}
	return string(symbols)
}

func checkHasCookie(r *http.Request) bool {
	cookie, err := r.Cookie("session_id")
	if err != nil || cookie.Valid() != nil {
		return false
	}
	return true
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	if !checkHasCookie(r) {
		w.Header().Set("Set-Cookie", fmt.Sprintf("session_id=%s; HttpOnly; Path=/", generateId()))
		fmt.Fprintf(w, "Welcome!")
		return
	}
	fmt.Fprintf(w, "Welcome back!")
}

func main() {
	http.Handle("/", http.HandlerFunc(setCookie))
	http.ListenAndServe(":8080", nil)
}
