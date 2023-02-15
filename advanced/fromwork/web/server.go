package web

type Server interface {
	Route(method string, pattern string, handeFunc func(c *Context))
	Start(address string) error
}
