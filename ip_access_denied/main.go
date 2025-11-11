package main

import (
	"fmt"
	"net/http"
)

func ipBlockerMiddleware(blockedIP string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		remoteIP := r.Header.Get("X-Real-IP") // strings.Split(r.RemoteAddr, ":")[0]
		if remoteIP == blockedIP {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func httpServing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Access granted")
}

func main() {
	http.Handle("/", ipBlockerMiddleware("192.168.0.1", http.HandlerFunc(httpServing)))
	http.ListenAndServe(":8080", nil)
}
