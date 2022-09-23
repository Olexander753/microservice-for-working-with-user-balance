package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Olexander753/microservice-for-working-with-user-balance/internal/config"
)

type Server struct {
	htppServer *http.Server
}

func (s *Server) Run(cfg *config.Config, handler http.Handler) error {
	log.Println("Run server")
	s.htppServer = &http.Server{
		Addr:           ":" + cfg.Server.Port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 28,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.htppServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.htppServer.Shutdown(ctx)
}
