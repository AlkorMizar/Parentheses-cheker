package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/AlkorMizar/Parentheses-cheker/internal/client"
	"github.com/AlkorMizar/Parentheses-cheker/internal/conf"
)

const (
	withSyncArgsLen  = 2
	withAsymcArgsLen = 3
)

var strLens = [...]int{2, 4, 8}

func main() {
	cfg, err := conf.NewConf()

	if err != nil {
		fmt.Println(err)
		return
	}

	var brClient *client.Client

	if len(os.Args) < withSyncArgsLen {
		fmt.Println("Client type did not chosen")
		return
	}

	if os.Args[1] == "sync" {
		brClient = client.NewClient(client.Sync, 0)
	} else {
		if len(os.Args) < withAsymcArgsLen {
			fmt.Println("No request limitation for async client")
			return
		}

		restrictNumber, err := strconv.Atoi(os.Args[2])

		if err != nil {
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
