package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	//initWeb()
	initWeb2()
}

func initWeb2() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", sayHello).Methods("post", "get").Host("127.0.0.1").Schemes("http")
	r.Handle("/hello2/{name}", &HelloWorldHandler{})
	r.HandleFunc("/hello3/{name:[a-z]+}", sayHello)
	// 路由前缀
	r.PathPrefix("hello4").HandlerFunc(sayHello)
	// 路由分组
	postRouter := r.PathPrefix("/post").Subrouter()
	postRouter.HandleFunc("", postIndex)
	postRouter.HandleFunc("/create", postCreate).Name("post.create")
	postRouter.HandleFunc("/delete", postCreate).Name("post.delete")

	fmt.Println(r.Get("post.create").URL())
	fmt.Println(r.Get("post.delete").URL())
	log.Fatal(http.ListenAndServe(":8089", r))
}

func postIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "首页")
}
func postCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "创建")
}

func initWeb() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", sayHello)
	r.Handle("/hello2/{name}", &HelloWorldHandler{})
	r.HandleFunc("/hello3/{name:[a-z]+}", sayHello)
	log.Fatal(http.ListenAndServe(":8089", r))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	// 接收参数
	params := mux.Vars(r)

	fmt.Fprintf(w, "hello,ycp!: %v", params)
}

type HelloWorldHandler struct {
}

func (receiver HelloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// 接收参数
	params := mux.Vars(r)

	fmt.Fprintf(w, "hello,ycp2222!: %v", params)
}
