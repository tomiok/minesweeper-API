package api

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/tomiok/minesweeper-API/internal/logs"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type server struct {
	*http.Server
}

func newServer(listening string, mux *chi.Mux) *server {
	s := &http.Server{
		Addr:         ":" + listening,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &server{s}
}

// Start runs ListenAndServe on the http.Server with graceful shutdown
func (srv *server) Start() {
	logs.Log().Info("starting server...")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logs.Sugar().Fatalf("could not listen on %s due to %s", srv.Addr, err.Error())
		}
	}()
	logs.Sugar().Infof("server is ready to handle requests %s", srv.Addr)
	srv.gracefulShutdown()
}

func (srv *server) gracefulShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	logs.Sugar().Infof("server is shutting down %s", sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	srv.SetKeepAlivesEnabled(false)
	if err := srv.Shutdown(ctx); err != nil {
		logs.Sugar().Fatalf("could not gracefully shutdown the server %s", err.Error())
	}
	logs.Log().Info("server stopped")
}
