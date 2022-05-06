package main

import (
	"fmt"

	"github.com/AlkorMizar/Parentheses-cheker/service"
)

func main() {
	err := service.LoadService()
	if err != nil {
		fmt.Print(err)
	}
}
