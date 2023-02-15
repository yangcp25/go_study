package web

import (
	"fmt"
	"net/http"
	"strings"
)

type Server interface {
	Route(method string, pattern string, handeFunc func(c *Context))
	Start(address string) error
}

type HttpSdk struct {
	Name        string
	HandlerFunc *Handler
}

type Handler struct {
	Handlers map[string]func(c *Context)
}

// Route 添加路由
func (h *HttpSdk) Route(method string, pattern string, handeFunc func(c *Context)) {
	key := getRouteKey(method, pattern)
	h.HandlerFunc.Handlers[key] = handeFunc
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
