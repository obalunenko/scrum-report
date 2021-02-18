// Package reporter provides functionality for report generation.
package reporter

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/obalunenko/scrum-report/internal/config"
)

// route Define a HTTP route with given logical name, http method, route pattern and handler function.
type route struct {
	description string
	method      string
	path        string
	handlerFunc http.HandlerFunc
}

// handler Describe all service API.
func routes() []route {
	return []route{
		{
			description: "index",
			method:      "GET",
			path:        "/",
			handlerFunc: indexHandler(),
		},
		{
			description: "create report",
			method:      "POST",
			path:        "/report",
			handlerFunc: createHandler(),
		},
	}
}

// Service holds all templates required by reporter.
type Service struct {
	config  config.Config
	handler http.Handler
}

// New creates new service from passed config.
func New(cfg config.Config) *Service {
	return &Service{
		config:  cfg,
		handler: newRouter(),
	}
}

// Run runs reporter service.
func (s *Service) Run() error {
	addr := fmt.Sprintf(":%s", s.config.Port)
	log.Debugf("address: %s", addr)

	return http.ListenAndServe(addr, s.handler)
}

// newRouter creates a new reporter service handler.
func newRouter() http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	// Register preflight handler
	options := http.HandlerFunc(optionsHandler)
	router.
		Methods(http.MethodOptions).
		Handler(options)

	for _, route := range routes() {
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
