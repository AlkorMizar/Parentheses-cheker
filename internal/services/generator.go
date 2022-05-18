package services

import (
	"net/http"
	"time"

	"github.com/AlkorMizar/Parentheses-cheker/internal/handlers"
	"github.com/AlkorMizar/Parentheses-cheker/internal/usecases"
)

// Route valid path for request
const Route = "/generate"
const Port = ":8080"

const (
	readTimeout  = 30
	writeTimeout = 90
	idleTimeout  = 120
)

// function called to load server
func Run() error {
	mux := http.NewServeMux()

	logic := usecases.NewBraces()
	h := handlers.NewHandlers(logic)

	mux.Handle(Route, h)

	s := http.Server{
		Addr:         Port,
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
