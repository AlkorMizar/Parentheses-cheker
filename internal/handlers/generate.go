package handlers

import (
	"log"
	"net/http"
	"strconv"
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
