package tcp

import (
	"context"
	"md5er/internal/servers/tcp/handlers"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func New() *Server {
	s := &Server{}
	routing := handlers.NewRouting()
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
