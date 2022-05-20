package main

import (
	"fmt"

	conf "github.com/AlkorMizar/Parentheses-cheker"
	"github.com/AlkorMizar/Parentheses-cheker/internal/services"
)

func main() {
	cfg, err := conf.NewConf()
	if err != nil {
		fmt.Print(err)
		return
	}

	service := services.NewService(*cfg)

	err = service.Run()
	if err != nil {
		fmt.Print(err)
	}
}
