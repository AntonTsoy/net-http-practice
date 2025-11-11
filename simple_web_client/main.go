package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type serverHandler struct{}

func (h *serverHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello from server")
}

func startServer(address string) {
	server := http.Server{
		Addr:    address,
		Handler: new(serverHandler),
	}

	//fmt.Println("Starting server at port 8080. Make a request on http://localhost:8080/")
	server.ListenAndServe()
}

func sendRequest(url string) (string, error) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	go startServer("localhost:8080")
	time.Sleep(time.Second)

	body, err := sendRequest("localhost:8080")
	fmt.Println(body, err)
}
