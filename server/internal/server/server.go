package server

import (
	"net/http"
	"rsoi-1/config"
	"time"
)

type Server interface {
	Run() error
}

type BaseServer struct {
	cfg *config.ServerConfig
}

func (s *BaseServer) createHttpServer() *http.Server {
	return &http.Server{
		Addr:         s.cfg.Host + ":" + s.cfg.Port,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
}
