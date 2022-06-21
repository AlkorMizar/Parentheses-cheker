package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/AlkorMizar/Parentheses-cheker/internal/client"
	"github.com/AlkorMizar/Parentheses-cheker/internal/conf"
)

var strLens = [...]int{2, 4, 8}

func main() {
	cfg, err := conf.NewConf()

	if err != nil {
		fmt.Println(err)
		return
	}

	var brClient *client.Client

	var isSync bool

	var restrictNumber int

	flag.BoolVar(&isSync, "isSync", true, "true -- sync, false -- async client")
	flag.IntVar(&restrictNumber, "restrNum", 1, "number of async request activeted through the task")

	flag.Parse()

	if isSync {
		brClient = client.NewClient(client.Sync, 0)
	} else {
		if restrictNumber <= 0 {
			fmt.Println("Incorrect request limitation for async client")
			return
		}

		brClient = client.NewClient(client.Async, restrictNumber)
	}

	url := "http://" + cfg.Address + cfg.Route + "?n="

	for _, l := range strLens {
		res, err := brClient.CalculateFor(url+strconv.Itoa(l), cfg.Cycles)
		if err != nil {
			fmt.Println(err)
			continue
		}

		res *= 100.0

		fmt.Printf("for length %d percentage = %.0f\n", l, res)
	}
}
