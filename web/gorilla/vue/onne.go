package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	initWeb()
}

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (s spaHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//panic("implement me")
	path, err := os.Getwd()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	path = filepath.Join(path, s.staticPath)
	//检查资源是否存在
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(writer, request, filepath.Join(s.staticPath, s.indexPath))
		return
	} else if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	http.FileServer(http.Dir(s.staticPath)).ServeHTTP(writer, request)
}

func initWeb() {
	mux := mux.NewRouter()
	spaH := &spaHandler{
		staticPath: "resource",
		indexPath:  "index.html",
	}
	mux.PathPrefix("/").Handler(spaH)

	srv := &http.Server{
		Handler: mux,
		Addr:    "127.0.0.1:8089",
		// 最佳实践：为服务器读写设置超时时间
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
