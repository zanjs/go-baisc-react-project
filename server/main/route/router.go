package route

import (
	"net/http"

	"../logger"
	"github.com/gorilla/mux"
)

//NewRouter create a configured route
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	// This will serve the react application under http://localhost:8000/
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("../../client/build")))

	return router
}
