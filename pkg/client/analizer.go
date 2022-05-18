package client

import (
	"errors"
	"io"
	"net/http"

	"github.com/AlkorMizar/Parentheses-cheker/internal/checker"
)

func CalculateFor(url string, cycles int) (float64, error) {
	return Calculate(getAllStr(url, cycles))
}

func getAllStr(url string, cycles int) (res []string) {
	res = make([]string, cycles)

	client := http.Client{}

	for i := 0; cycles > 0; cycles-- {
		body, err := getBody(client, url)
		if err == nil {
			res[i] = body
			i++
		}
	}

	return res
}

func getBody(cl http.Client, url string) (string, error) {
	resp, err := cl.Get(url)

	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusAccepted {
		return "", err
	}

	body, errBody := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if errBody != nil {
		return "", err
	}

	return string(body), nil
}

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