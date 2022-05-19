package client

import (
	"errors"

	"github.com/AlkorMizar/Parentheses-cheker/internal/checker"
)

type Client struct {
	getAllStr      func(url string, cycles int, restrictNumber int) []string
	restrictNumber int
}

func NewClient(getter func(url string, cycles int, restrictNumber int) []string, restrNum int) *Client {
	return &Client{
		getAllStr:      getter,
		restrictNumber: restrNum,
	}
}

func (cl *Client) CalculateFor(url string, cycles int) (float64, error) {
	return Calculate(cl.getAllStr(url, cycles, cl.restrictNumber))
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
