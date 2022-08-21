package handlers

import (
	"io"
	"net/http"
)

func Post(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "welcome to post")
}
