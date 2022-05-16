package main

import (
	"fmt"
	"strconv"

	"github.com/AlkorMizar/Parentheses-cheker/pkg/client"
	"github.com/AlkorMizar/Parentheses-cheker/pkg/services"
)

const N = 1000

var strLens = [...]int{2, 4, 8}
var url = "http://localhost" + services.Port + services.ServiceRoute + "?n="

func main() {
	for _, l := range strLens {
		res, err := client.CalculateFor(url+strconv.Itoa(l), N)
		if err != nil {
			fmt.Println(err)
			continue
		}

		res *= 100.0

		fmt.Printf("for length %d percentage = %.0f\n", l, res)
	}
}
