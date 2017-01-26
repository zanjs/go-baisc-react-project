package route

import (
	"net/http"

	"../handler"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"TodoIndex",
		"GET",
		"/services/todos",
		handler.TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/services/todos/{todoId}",
		handler.TodoShow,
	},
	Route{
		"TodoCreate",
		"POST",
		"/services/todos",
		handler.TodoCreate,
	},
}
