package server

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/obalunenko/logger"
)

type Server struct {
	name    string
	ctx     context.Context
	wg      *sync.WaitGroup
	srv     *http.Server
	errChan chan error
}

type ShutdownFunc func(wg *sync.WaitGroup, s *http.Server)

func NewServer(ctx context.Context, wg *sync.WaitGroup, name string, port string, logWriter io.Writer,
	handler http.Handler, shutdownFunc ShutdownFunc) *Server {
	errLog := log.New(logWriter, fmt.Sprintf("%s: ", name), log.LstdFlags)

	srv := http.Server{
		Addr:              net.JoinHostPort("", port),
		Handler:           handler,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          errLog,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	srv.RegisterOnShutdown(func() {
		shutdownFunc(wg, &srv)
	})

	return &Server{
		name:    name,
		ctx:     ctx,
		wg:      wg,
		srv:     &srv,
		errChan: make(chan error, 1),
	}
}

func (s *Server) Errors() <-chan error {
	return s.errChan
}

func (s *Server) Run() {
	go s.startServer()
	go s.handleShutdown()

	logger.WithFields(s.ctx, logger.Fields{
		"addr": s.srv.Addr,
		"name": s.name,
	}).Info("Up and running")
}

func (s *Server) startServer() {
	if err := s.srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			s.errChan <- fmt.Errorf("%s[%s]: srv error: %w", s.name, s.srv.Addr, err)
		}
	}
}

func (s *Server) handleShutdown() {
	defer s.wg.Done()

	<-s.ctx.Done()

	s.srv.ErrorLog.Print("shutting down")

	if err := s.srv.Shutdown(s.ctx); err != nil && !errors.Is(err, context.Canceled) {
		s.srv.ErrorLog.Printf("shutdouwn error: %v \n", err)
	}
}
