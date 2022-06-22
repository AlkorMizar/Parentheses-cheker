package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/AlkorMizar/Parentheses-cheker/internal/client"
	"github.com/AlkorMizar/Parentheses-cheker/internal/conf"
)

var strLens = [...]int{2, 4, 8}

func getType(s string) (client.Type, error) {
	if s == "sync" {
		return client.Sync, nil
	} else if s == "async" {
		return client.Async, nil
	}

	return client.Sync, fmt.Errorf(`client type either "sync" or "async" `)
}

func getRestrNum(s string) (int, error) {
	restN, err := strconv.Atoi(s)
	if err != nil {
		return 1, err
	}

	if restN < 0 {
		return 1, fmt.Errorf("restrNum > 0")
	}

	return restN, err
}

func main() {
	cfg, err := conf.NewConf()

	if err != nil {
		fmt.Println(err)
		return
	}

	var brClient *client.Client

	clType := client.Sync

	restrictNumber := 1

	flag.Func("type", "sync/async client", func(s string) error {
		clType, err = getType(s)
		return err
	})

	flag.Func("restrNum", "number of async request activeted through the task", func(s string) error {
		restrictNumber, err = getRestrNum(s)
		return err
	})

	flag.Parse()

	brClient = client.NewClient(clType, restrictNumber)

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
