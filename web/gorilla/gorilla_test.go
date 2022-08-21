package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testIntiWeb6(t *testing.T) {
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
