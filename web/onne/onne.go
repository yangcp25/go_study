package main

import (
	"fmt"
	"net/http"
)

func main() {
	//initWeb()
	//initWebRoute()
	initWebRoute2()
}

func initWebRoute2() {
	hander := Test2{}
	test := Test{}

	server := http.Server{Addr: ":8089"}
	http.Handle("/", &hander)
	http.Handle("/test", &test)
	http.Handle("/hander", &hander)

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("发生错误!err:", err)
	}
}

func initWebRoute() {
	hander := MyHandler{}
	err := http.ListenAndServe(":8089", &hander)

	if err != nil {
		fmt.Println("发生错误!err:", err)
	}
}

type MyHandler struct {
}

type Test struct {
}

func (t *Test) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sayHello2(w, r)
	return
}

type Test2 struct {
}

func (t *Test2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sayHello(w, r)
	return
}

func sayHello2(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("url:", r.URL.Path)
	fmt.Println("Scheme", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Printf("key:%v;val:%v", k, v)
	}
	fmt.Fprintf(w, "你隔壁11111111111111")
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHello(w, r)
		return
	}
	http.NotFound(w, r)
	return
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
