package handlers

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func User(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "welcome to user")
	params := mux.Vars(request)
	id := params["id"]
	io.WriteString(writer, "id:"+id)
}
