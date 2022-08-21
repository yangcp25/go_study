package handlers

import (
	"io"
	"net/http"
)

func Home(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "welcome to ycp home!")
}
