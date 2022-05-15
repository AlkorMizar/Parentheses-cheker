package main

import (
	"fmt"
	"strconv"

	"github.com/AlkorMizar/Parentheses-cheker/analizer"
	"github.com/AlkorMizar/Parentheses-cheker/service"
)

const N = 1000

var strLens = [...]int{2, 4, 8}
var url = "http://localhost" + service.Port + service.ServiceRoute + "?n="

func main() {
	for _, l := range strLens {
		res, err := analizer.CalculateFor(url+strconv.Itoa(l), N)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("for length %d percentage = %.0f\n", l, res*100)
	}
}
