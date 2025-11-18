package api

import (
	"context"
	"gateway"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	config "github.com/mwdev22/gocfg"
	"github.com/rs/cors"
)

type Api struct {
	cfg         *config.Config
	matchmaking gateway.MatchmakingClient
	session     gateway.SessionClient
	player      gateway.PlayerClient
}

func New(cfg *config.Config, opts ...func(*Api)) *Api {
	api := &Api{
		cfg: cfg,
	}
	for _, opt := range opts {
		opt(api)
	}
	return api
}

func WithMatchmakingClient(cli gateway.MatchmakingClient) func(*Api) {
	return func(a *Api) {
		a.matchmaking = cli
	}
}

func WithSessionClient(cli gateway.SessionClient) func(*Api) {
	return func(a *Api) {
		a.session = cli
	}
}

func WithPlayerClient(cli gateway.PlayerClient) func(*Api) {
	return func(a *Api) {
		a.player = cli
	}
}

func (a *Api) Run() error {
	mux := chi.NewMux()

	a.Mount(mux)

	c := cors.New(cors.Options{
		AllowedOrigins:      []string{"*"},
		AllowCredentials:    true,
		AllowedMethods:      []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowPrivateNetwork: true,
		AllowedHeaders:      []string{"*"},
	})

	server := &http.Server{
		Addr:         a.cfg.Addr,
		Handler:      c.Handler(mux),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("Server is listening on %s", a.cfg.Addr)
		serverErrors <- server.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		return err
	case sig := <-quit:
		log.Printf("Received signal %s, shutting down gracefully...", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Graceful shutdown failed: %v", err)
			return err
		}
		log.Println("Server stopped gracefully")
		return nil
	}
}
