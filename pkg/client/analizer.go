package client

import (
	"errors"
	"io"
	"net/http"

	"github.com/AlkorMizar/Parentheses-cheker/internal/checker"
)

func CalculateFor(url string, cycles int) (float64, error) {
	client := http.Client{}
	all := 0.0
	checked := 0.0
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
		return -1, errors.New("service not working")
	}
	return checked / all, nil
}
