package main

import (
	"fmt"
	"mosaic2/tool"
	"net/http"
)

// 马赛克程序
func main() {
	mux := http.NewServeMux()

	files := http.FileServer(http.Dir("public"))

	// 路由
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", tool.Upload)
	mux.HandleFunc("/mosaic", tool.Mosaic)

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	tool.TILESDB = tool.TilesDb()
	fmt.Println("马赛克服务器已启动...")
	server.ListenAndServe()
}
