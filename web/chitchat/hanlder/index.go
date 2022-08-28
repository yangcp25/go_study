package hanlder

import (
	"chichat/model"
	"fmt"
	"html/template"
	"net/http"
)

// 论坛首页路由处理器方法
func Index(w http.ResponseWriter, r *http.Request) {
	files := []string{"views/layout.html", "views/navbar.html", "views/index.html"}
	templates := template.Must(template.ParseFiles(files...))
	threads, err := model.Threads()
	fmt.Println(threads)
	if err == nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}
