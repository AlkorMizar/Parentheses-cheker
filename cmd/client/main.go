package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"

	conf "github.com/AlkorMizar/Parentheses-cheker"
	"github.com/AlkorMizar/Parentheses-cheker/internal/client"
)

var strLens = [...]int{2, 4, 8}

func main() {
	cfg, err := conf.NewConf()

	if err != nil {
		fmt.Println(err)
		return
	}

	url := cfg.Address + cfg.Port + cfg.Route + "?n="

	var brClient *client.Client
	if os.Args[1] == "sync" {
		brClient = client.NewClient(fetchAllStrSync, 0)
	} else {
		brClient = client.NewClient(fetchAllStrAsync, 0)
	}

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

func fetchAllStrSync(url string, cycles int, restrictNumber int) (res []string) {
	res = make([]string, cycles)

	client := http.Client{}

	for i := 0; cycles > 0; cycles-- {
		body, err := getBody(client, url)
		if err != nil {
			continue
		}

		res[i] = body
		i++

	}

	return res
}

func fetchAllStrAsync(url string, cycles int, restrictNumber int) (res []string) {
	var mutx sync.Mutex
	res = make([]string, cycles)
	client := http.Client{}
	restr := make(chan struct{}, restrictNumber)

	for i := 0; cycles > 0; cycles-- {
		go func() {
			defer func() {
				<-restr
			}()
			body, err := getBody(client, url)
			if err != nil {
				return
			}

			mutx.Lock()

			res[i] = body
			i++

			mutx.Unlock()

		}()

		restr <- struct{}{}
	}

	return res
}

func getBody(cl http.Client, url string) (string, error) {
	resp, err := cl.Get(url)

	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("status not OK")
	}

	body, errBody := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if errBody != nil {
		return "", err
	}

	return string(body), nil
}
