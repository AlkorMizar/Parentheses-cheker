package client

import (
	"errors"
	"io"
	"net/http"
	"sync"

	"github.com/AlkorMizar/Parentheses-cheker/internal/checker"
)

type Client struct {
	fetchAllStr func(url string, cycles int) []string
}

type Type int

const Sync Type = 1
const Async Type = 2

// function called to create client that will use given function to fetch array of strings and then analize given strings
func NewClient(tp Type, limiteNum int) *Client {
	var fetcher func(string, int) []string
	switch tp {
	case Sync:
		{
			fetcher = fetchAllStrSync
		}
	case Async:
		{
			fetcher = func(url string, cycles int) []string {
				return fetchAllStrAsync(url, cycles, limiteNum)
			}
		}
	}

	return &Client{
		fetchAllStr: fetcher,
	}
}

// function that user can call
func (cl *Client) CalculateFor(url string, cycles int) (float64, error) {
	return Calculate(cl.fetchAllStr(url, cycles))
}

func fetchAllStrSync(url string, cycles int) (res []string) {
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
	var wg sync.WaitGroup
	res = make([]string, cycles)
	client := http.Client{}
	restr := make(chan struct{}, restrictNumber)

	for i := 0; cycles > 0; cycles-- {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
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
	wg.Wait()
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

// function that analize strings
func Calculate(inp []string) (float64, error) {
	all := 0.0
	checked := 0.0

	for _, v := range inp {
		if checker.Check(v) {
			checked++
		}
		all++
	}

	if all == 0 {
		return -1, errors.New("service not working")
	}

	return checked / all, nil
}
