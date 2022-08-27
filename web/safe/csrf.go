package main

import (
	md52 "crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"html/template"
	"io"
	"net/http"
	"strconv"
)

func main() {
	//initCsrf()
	initApiCsrf()
}

type User struct {
	Id   int
	Name string
}

func initApiCsrf() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.Use(csrf.Protect([]byte(generateMd5("123456"))))

	api.HandleFunc("/user/{id}", handUser)
	http.ListenAndServe(":8089", r)
}

func handUser(writer http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)

	id2, _ := strconv.Atoi(params["id"])
	user := User{
		Id:   id2,
		Name: "ycp",
	}
	writer.Header().Set("X-Csrf-Token", csrf.Token(request))
	res, err := json.Marshal(user)

	if err != nil {
		panic(err)
	}
	writer.Write(res)
}

func initCsrf() {
	c := mux.NewRouter()

	c.HandleFunc("/signup", handSignUp)

	http.ListenAndServe(":8089",
		csrf.Protect([]byte(generateMd5("123456")), csrf.Secure(false))(c))
}

func generateMd5(code string) string {
	md5 := md52.New()
	_, err := io.WriteString(md5, code)
	if err != nil {
		return ""
	}

	res := md52.Sum(nil)

	return hex.EncodeToString(res[:16])
}

func handSignUp(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("html/test.html"))
	t.Execute(writer, map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(request),
	})
	// 我们还可以通过 csrf.Token(r) 直接获取令牌并将其设置到请求头：w.Header.Set("X-CSRF-Token", token)
	// 这在发送 JSON 响应到客户端或者前端 JavaScript 框架时很有用
}
