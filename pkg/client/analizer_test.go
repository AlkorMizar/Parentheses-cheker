package client_test

import (
	"testing"

	"github.com/AlkorMizar/Parentheses-cheker/pkg/client"
)

func TestCalculate(t *testing.T) {
	tests := map[string]struct {
		input  []string
		output float64
	}{
		"simple": {
			input:  []string{"()", "[]", "{}"},
			output: 1.0,
		},
		"has unbalanced": {
			input:  []string{"()[]", "}{[]", "{})("},
			output: 1.0 / 3.0,
		},
		"all unbalanced": {
			input:  []string{"[]]{}]", "}{[]", "{})(", "{(){}"},
			output: 0.0,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := client.Calculate(tc.input)

			if err != nil || got != tc.output {
				t.Fatalf("erorr in test %s, expected: %v, got: %v", name, tc.output, got)
			}
		})
	}
}
