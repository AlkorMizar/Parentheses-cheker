package service

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ServiceRoute valid path for request
const ServiceRoute = "/generate"

const (
	readTimeout  = 30
	writeTimeout = 90
	idleTimeout  = 120
)

type businessL func(leng int) string

type Handlers struct {
	businessLogic businessL
}

func NewHandlers(funct businessL) *Handlers {
	return &Handlers{
		businessLogic: funct,
	}
}

// function called to load server
func LoadService() error {
	mux := http.NewServeMux()

	h := NewHandlers(Generate)

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

	_, err = w.Write([]byte(h.businessLogic(n)))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	}
}

// business logic that generates strings with braces
func Generate(leng int) string {
	var builder strings.Builder

	braces := []rune{'(', ')', '{', '}', '[', ']'}
	amount := len(braces)

	for leng > 0 {
		builder.WriteRune(braces[rand.Int31n(int32(amount))]) //nolint:gosec // this is used to simplify programm
		leng--
	}

	return builder.String()
}
