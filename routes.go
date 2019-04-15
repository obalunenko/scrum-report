package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// route Define a HTTP route with given logical name, http method, route pattern and handler function
type route struct {
	description string
	method      string
	path        string
	handlerFunc http.HandlerFunc
}

// routes Describe all service API

var routes = []route{
	{
		description: "index",
		method:      "GET",
		path:        "/",
		handlerFunc: indexHandler,
	},
	{
		description: "create report",
		method:      "POST",
		path:        "/report",
		handlerFunc: createHandler,
	},
}

// NewRouter creates a new router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Register preflight handler
	options := http.HandlerFunc(optionsHandler)
	router.
		Methods(http.MethodOptions).
		Handler(options)

	for _, route := range routes {

		handler := http.Handler(route.handlerFunc)
		handler = loggerHandler(handler, route.description)
		handler = handlers.CompressHandler(handler)

		router.
			Methods(route.method).
			Path(route.path).
			Name(route.description).
			Handler(handler)
	}

	return router
}
