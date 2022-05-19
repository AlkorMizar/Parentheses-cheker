package services

import (
	"net/http"

	conf "github.com/AlkorMizar/Parentheses-cheker"
	"github.com/AlkorMizar/Parentheses-cheker/internal/handlers"
	"github.com/AlkorMizar/Parentheses-cheker/internal/usecases"
)

// function called to load server
func Run(c *conf.Config) error {
	mux := http.NewServeMux()

	logic := usecases.NewBraces()
	h := handlers.NewHandlers(logic)

	mux.Handle(c.Route, h)

	s := http.Server{
		Addr:    c.Port,
		Handler: mux,
	}

	err := s.ListenAndServe()

	if err != nil {
		return err
	}

	return nil
}
