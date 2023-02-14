package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		str := []byte{'h', 'e', 'l', 'l', '0'}
		writer.Write(str)
	})
	http.ListenAndServe("localhost:8081", nil)
}
