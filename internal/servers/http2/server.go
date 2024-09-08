package http2

import (
	"context"
	"md5er/internal/md5"
	"md5er/internal/servers/http2/handlers"
	"net/http"
	"time"
)

type Server struct {
	server     *http.Server
	ms5service *md5.MD5service
}

func New(md5service *md5.MD5service) *Server {
	s := &Server{
		ms5service: md5service,
	}
	routing := handlers.NewRouting(s.ms5service)
	s.server = &http.Server{
		Addr:         "127.0.0.1:8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      routing.Router,
	}
	return s
}
func (s Server) Start() error {
	return s.server.ListenAndServe()
}

func (s Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
