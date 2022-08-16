package main

import (
	"fmt"
	"net/http"
)

func main() {
	initWeb()
}

func initWeb() {
	http.HandleFunc("/", sayHello)

	err := http.ListenAndServe(":8089", nil)

	if err != nil {
		fmt.Println("发生错误!err:", err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("url:", r.URL.Path)
	fmt.Println("Scheme", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Printf("key:%v;val:%v", k, v)
	}
	fmt.Println("你好。ycp")
	msg := []byte("你好的很")
	w.Write(msg)
	fmt.Fprintf(w, "你隔壁")
}
