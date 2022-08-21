package routes

import (
	"net/http"
	. "routes/handlers"
)

type WebRoute struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type WebRoutes []WebRoute

var webRoutes = WebRoutes{
	WebRoute{
		"Home",
		"Get",
		"/",
		Home,
	}, WebRoute{
		"Post",
		"Get",
		"/posts",
		Post,
	}, WebRoute{
		"User",
		"Get",
		"/user/{id}",
		User,
	},
}
