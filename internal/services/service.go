package services

import (
	"net/http"

	conf "github.com/AlkorMizar/Parentheses-cheker"
	"github.com/AlkorMizar/Parentheses-cheker/internal/braces"
	"github.com/AlkorMizar/Parentheses-cheker/internal/handlers"
)

type Service struct {
	service *http.Server
}

// function called to create service and configure it
func NewService(c conf.Config) *Service {
	mux := http.NewServeMux()

	logic := braces.NewBraces()
	h := handlers.NewHandlers(logic)

	mux.Handle(c.Route, h)

	server := http.Server{
		Addr:    c.Port,
		Handler: mux,
	}

	return &Service{
		service: &server,
	}
}

// function called to run service
func (s *Service) Run() error {
	err := s.service.ListenAndServe()

	if err != nil {
		return err
	}

	return nil
}
