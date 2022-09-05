package handlers

import (
	"chichat/models"
	"net/http"
)

// 论坛首页路由处理器方法
func Index(w http.ResponseWriter, r *http.Request) {
	threads, _ := models.Threads()
	generateHTML(w, threads, "layout", "navbar", "index")
}
