package web

import (
	"fmt"
	"net/http"
	"strings"
)

type IRoutable interface {
	Route(method string, pattern string, handeFunc func(c *Context))
}

type IHandler interface {
	http.Handler
	IRoutable
}

type Server interface {
	IRoutable
	Start(address string) error
}

type HttpSdk struct {
	Name        string
	HandlerFunc IHandler
}

// 确保一定实现了接口
var _ IHandler = &Handler{}

type Handler struct {
	Handlers map[string]func(c *Context)
}

func (h Handler) Route(method string, pattern string, handeFunc func(c *Context)) {
	key := getRouteKey(method, pattern)
	h.Handlers[key] = handeFunc
}

// Route 添加路由
func (h *HttpSdk) Route(method string, pattern string, handeFunc func(c *Context)) {
	h.HandlerFunc.Route(method, pattern, handeFunc)
}

func (h *HttpSdk) Start(address string) error {
	err := http.ListenAndServe(address, h.HandlerFunc)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key := getRouteKey(strings.ToLower(request.Method), request.URL.Path)
	if handler, ok := h.Handlers[key]; ok {
		c := &Context{request, writer}
		handler(c)
	} else {
		fmt.Println("not found")
	}
}

func getRouteKey(method string, path string) string {
	return fmt.Sprintf("%s#%s", method, path)
}
