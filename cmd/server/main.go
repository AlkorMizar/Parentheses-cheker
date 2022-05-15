package main

import (
	"fmt"

	"github.com/AlkorMizar/Parentheses-cheker/pkg/services"
)

func main() {
	err := services.LoadService()
	if err != nil {
		fmt.Print(err)
	}
}
