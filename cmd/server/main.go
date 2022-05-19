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
	err = services.Run(cfg)
	if err != nil {
		fmt.Print(err)
	}
}
