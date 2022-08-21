package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func main() {
	//initWeb()
	//initWeb2()
	// 路由中间件
	//initWeb3()
	// 处理静态资源
	//initWeb4()
	// 处理跨域请求
	//initWeb5()
	// 健康检查
	//intiWeb6()
}

func intiWeb6(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/health", handHealth)
	//r.Use(mux.CORSMethodMiddleware(r))
	//http.ListenAndServe(":8091", r)

	// http接口进行测试
	req, _ := http.NewRequest("GET", "/health", nil)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// 检测状态码
	if rr.Code != http.StatusOK {
		t.Error("not ok")
	}

	// 健康检查结果

	if rr.Body.String() != "alive" {
		t.Error("not alive")
	}
}

func handHealth(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	io.WriteString(writer, "alive")
}

func initWeb5() {
	r := mux.NewRouter()
	r.HandleFunc("/cors", handOption).Methods(http.MethodGet, http.MethodPut, http.MethodOptions)
	r.Use(mux.CORSMethodMiddleware(r))
	http.ListenAndServe(":8091", r)
}

func handOption(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set("Access-Control-Allow-Methods", "*")

	if r.Method == http.MethodOptions {
		return
	}
	writer.Write([]byte("ycp"))
}

func initWeb4() {
	mux := mux.NewRouter()
	mux.HandleFunc("/index", handIndex)
	mux.Use(logMiddleware)

	var dir string = "static"
	//flag.StringVar(&dir, "dir", ".", "静态资源目录，默认为当前目录")
	//flag.Parse()

	mux.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir(dir))))
	// 应用于子路由
	http.ListenAndServe(":8089", mux)
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println("log...")
		next.ServeHTTP(writer, request)
	})
}
func postLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		token := request.FormValue("token")
		if token == "ycp" {
			log.Println("success token")
			next.ServeHTTP(writer, request)
		} else {
			http.Error(writer, "err token", http.StatusForbidden)
		}
	})
}
func initWeb3() {
	mux := mux.NewRouter()
	mux.HandleFunc("/index", handIndex)
	mux.Use(logMiddleware)
	// 应用于子路由
	prePost := mux.PathPrefix("/post").Subrouter()
	prePost.HandleFunc("", handPost)
	prePost.Use(postLogMiddleware)
	http.ListenAndServe(":8089", mux)
}

func handIndex(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "index")
}

func handPost(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "post")
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
