package main

import (
	"fmt"
	"strconv"

	conf "github.com/AlkorMizar/Parentheses-cheker"
	"github.com/AlkorMizar/Parentheses-cheker/pkg/client"
)

var strLens = [...]int{2, 4, 8}

func main() {
	cfg, err := conf.NewConf()

	if err != nil {
		fmt.Println(err)
		return
	}

	url := cfg.Address + cfg.Port + cfg.Route + "?n="

	for _, l := range strLens {
		res, err := client.CalculateFor(url+strconv.Itoa(l), cfg.Cycles)
		if err != nil {
			fmt.Println(err)
			continue
		}

		res *= 100.0

		fmt.Printf("for length %d percentage = %.0f\n", l, res)
	}
}
