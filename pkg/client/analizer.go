package client

import (
	"errors"
	"io"
	"net/http"

	"github.com/AlkorMizar/Parentheses-cheker/internal/checker"
)

func CalculateFor(url string, cycles int) (int, error) {
	client := http.Client{}
	all := 0
	checked := 0
	for ; cycles > 0; cycles-- {
		resp, err := client.Get(url)
		if err != nil {
			continue
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			continue
		}
		if checker.Check(string(body)) {
			checked++
		}
		all++
	}
	if all == 0 {
		return -1, errors.New("Service not working")
	}
	return checked / all, nil
}
