package checker_test

import (
	"testing"

	checker "github.com/AlkorMizar/Parentheses-cheker"
)

func TestCheck(t *testing.T) {
	tests := map[string]struct {
		input  string
		output bool
	}{
		"simple": {
			input:  "()",
			output: true,
		},
		"empty": {
			input:  "",
			output: true,
		},
		"simple unbalanced": {
			input:  ")(",
			output: false,
		},
		"different braces": {
			input:  "[{}()](){[]}",
			output: true,
		},
		"different braces unbalanced": {
			input:  "[)(}{()}[}]{{{",
			output: false,
		},
		"with other symbols": {
			input:  "((1 + 2) * 3) - 4)/5+[dasd]-{a(asd)}",
			output: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := checker.Check(tc.input)

			if got != tc.output {
				t.Fatalf("erorr in test %s, expected: %v, got: %v", name, tc.output, got)
			}
		})
	}
}
