package services

import (
	"net/http"
	"time"

	"github.com/AlkorMizar/Parentheses-cheker/internal/handlers"
	"github.com/AlkorMizar/Parentheses-cheker/internal/usecases"
)

// ServiceRoute valid path for request
const ServiceRoute = "/generate"

const (
	readTimeout  = 30
	writeTimeout = 90
	idleTimeout  = 120
)

// function called to load server
func LoadService() error {
	mux := http.NewServeMux()

	logic := usecases.NewBraces()
	h := handlers.NewHandlers(logic)

	mux.Handle(ServiceRoute, h)

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  readTimeout * time.Second,
		WriteTimeout: writeTimeout * time.Second,
		IdleTimeout:  idleTimeout * time.Second,
		Handler:      mux,
	}

	err := s.ListenAndServe()

	if err != nil {
		return err
	}

	return nil
}
