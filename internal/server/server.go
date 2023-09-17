package server

import (
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(port string, handler http.Handler) *Server {
	return &Server{httpServer: &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}
