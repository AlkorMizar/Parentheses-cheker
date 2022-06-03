package server

import (
	"net/http"
)

type Server struct {
	server *http.Server
}

// function called to create service and configure it
func NewServer(Addr string, mux *http.ServeMux) *Server {

	server := http.Server{
		Addr:    Addr,
		Handler: mux,
	}

	return &Server{
		server: &server,
	}
}

// function called to run service
func (s *Server) Run() error {
	err := s.server.ListenAndServe()

	if err != nil {
		return err
	}

	return nil
}
