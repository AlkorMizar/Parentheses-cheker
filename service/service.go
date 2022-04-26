package service

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

// ServiceRoute valid path for request
const ServiceRoute = "/generate"

const (
	readTimeout  = 30
	writeTimeout = 90
	idleTimeout  = 120
)

// function called to load server
func LoadService() {
	mux := http.NewServeMux()

	mux.HandleFunc(ServiceRoute, GenrateHandler)

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  readTimeout * time.Second,
		WriteTimeout: writeTimeout * time.Second,
		IdleTimeout:  idleTimeout * time.Second,
		Handler:      mux,
	}

	err := s.ListenAndServe()

	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}

// function that handles request
func GenrateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	n, err := strconv.Atoi(r.URL.Query().Get("n"))

	if err != nil || n < 0 {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	_, err = w.Write([]byte(generate(n)))

	if err != nil {
		log.Print(err)
	}
}

// business logic that generates strings with braces
func generate(leng int) string {
	return ""
}
