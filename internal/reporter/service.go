// Package reporter provides functionality for report generation.
package reporter

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/obalunenko/logger"

	"github.com/obalunenko/scrum-report/internal/server"
)

// route Define an HTTP route with given logical name, http method, route pattern and handler function.
type route struct {
	description string
	method      string
	path        string
	handlerFunc http.HandlerFunc
}

// handler Describe all service API.
func routes(ctx context.Context) []route {
	return []route{
		{
			description: "index",
			method:      "GET",
			path:        "/",
			handlerFunc: indexHandler(ctx),
		},
		{
			description: "create report",
			method:      "POST",
			path:        "/report",
			handlerFunc: createHandler(ctx),
		},
	}
}

// Service represents reporter service instance.
type Service struct {
	appServer *server.Server
	wg        *sync.WaitGroup
	stopChan  chan os.Signal
	ctx       struct {
		val        context.Context
		cancelFunc context.CancelFunc
	}
}

// Params holds Service create parameters.
type Params struct {
	AppName string
	Port    string
}

// New creates new service from passed config.
func New(ctx context.Context, params Params) *Service {
	ctx, cancel := context.WithCancel(ctx)

	handler := newRouter(ctx)

	var wg sync.WaitGroup

	wg.Add(1)

	logWriter := log.FromContext(ctx).Writer()

	srv := server.New(
		ctx,
		&wg,
		params.AppName,
		params.Port,
		logWriter,
		handler,
		func(wg *sync.WaitGroup, s *http.Server) {
			defer wg.Done()

			s.ErrorLog.Println("Disable keep-alive")

			s.SetKeepAlivesEnabled(false)

			if err := logWriter.Close(); err != nil {
				s.ErrorLog.Printf("failed to close log writer: %v", err)
			}
		},
	)

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT)

	return &Service{
		appServer: srv,
		wg:        &wg,
		stopChan:  stopChan,
		ctx: struct {
			val        context.Context
			cancelFunc context.CancelFunc
		}{
			val:        ctx,
			cancelFunc: cancel,
		},
	}
}

// Run runs Service and returns channel that will indicate when Service is finished execution.
func (s *Service) Run() chan struct{} {
	s.wg.Add(1)

	go s.appServer.Run()

	doneChan := make(chan struct{})

	go func() {
	loop:
		for {
			select {
			case sig := <-s.stopChan:
				log.WithField(s.ctx.val, "signal", sig.String()).Warn("Signal received")

				break loop
			case err := <-s.appServer.Errors():
				if err != nil {
					log.WithError(s.ctx.val, err).Error("server error")

					break loop
				}
			}
		}

		s.ctx.cancelFunc()

		s.wg.Wait()

		close(doneChan)
	}()

	return doneChan
}

// newRouter creates a new reporter service handler.
func newRouter(ctx context.Context) http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	// Register preflight handler
	options := http.HandlerFunc(optionsHandler)
	router.
		Methods(http.MethodOptions).
		Handler(options)

	for _, route := range routes(ctx) {
		handler := http.Handler(route.handlerFunc)
		handler = loggerHandler(ctx, handler, route.description)
		handler = handlers.CompressHandler(handler)

		router.
			Methods(route.method).
			Path(route.path).
			Name(route.description).
			Handler(handler)
	}

	return router
}
