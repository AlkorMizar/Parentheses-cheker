package service

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/AlkorMizar/Parentheses-cheker/service/usecases"
)

// ServiceRoute valid path for request
const ServiceRoute = "/generate"
const Port = ":8080"

const (
	readTimeout  = 30
	writeTimeout = 90
	idleTimeout  = 120
)

type braces interface {
	Generate(int) string
}

type Handlers struct {
	bracesLogic braces
}

func NewHandlers(brL braces) *Handlers {
	return &Handlers{
		bracesLogic: brL,
	}
}

// function called to load server
func LoadService() error {
	mux := http.NewServeMux()

	logic := usecases.NewBraces()
	h := NewHandlers(logic)

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

// function that handles request
func (h *Handlers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	n, err := strconv.Atoi(r.URL.Query().Get("n"))

	if err != nil || n < 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = w.Write([]byte(h.bracesLogic.Generate(n)))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	}
}
