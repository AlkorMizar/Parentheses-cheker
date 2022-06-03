package main

import (
	"fmt"
	"net/http"

	"github.com/AlkorMizar/Parentheses-cheker/internal/braces"
	"github.com/AlkorMizar/Parentheses-cheker/internal/conf"
	"github.com/AlkorMizar/Parentheses-cheker/internal/handlers"
	"github.com/AlkorMizar/Parentheses-cheker/internal/server"
)

func main() {
	cfg, err := conf.NewConf()
	if err != nil {
		fmt.Print(err)
		return
	}

	mux := http.NewServeMux()

	logic := braces.NewBraces()
	h := handlers.NewHandlers(logic)

	mux.Handle(cfg.Route, h)

	service := server.NewServer(cfg.Address, mux)

	err = service.Run()
	if err != nil {
		fmt.Print(err)
	}
}
