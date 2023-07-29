package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type Middleware func(Handler) Handler

func Logging(h Handler) Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("logging")
		h.ServeHTTP(w, r)
	})
}

func Authentication(h Handler) Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("authenticating")
		h.ServeHTTP(w, r)
	})
}

func handler() Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handler")
		w.Write([]byte("OK"))
	})
}

func main() {
	finalHandler := Logging(Authentication(handler()))

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()

	finalHandler.ServeHTTP(w, req)
}
