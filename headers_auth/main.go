package main

import (
	"fmt"
	"net/http"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func httpServing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Authorized access")
}

func main() {
	http.Handle("/", authMiddleware(http.HandlerFunc(httpServing)))
	http.ListenAndServe(":8080", nil)
}
