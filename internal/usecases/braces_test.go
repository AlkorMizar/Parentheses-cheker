package usecases_test

import (
	"regexp"
	"testing"

	"github.com/AlkorMizar/Parentheses-cheker/internal/usecases"
)

func TestGenerateResult(t *testing.T) {
	tests := map[string]struct {
		lenIn  int
		lenOut int
	}{
		"correct length": {
			lenIn:  8,
			lenOut: 8,
		},
		"odd lengrh": {
			lenIn:  7,
			lenOut: 7,
		},
		"big length": {
			lenIn:  100,
			lenOut: 100,
		},
		"zero": {
			lenIn:  0,
			lenOut: 0,
		},
	}

	for name, tc := range tests {
		generator := usecases.NewBraces()
		t.Run(name, func(t *testing.T) {
			got := generator.Generate(tc.lenIn)
			re := regexp.MustCompile("[^(){}\\[\\]]+") //nolint:gosimple // this is the only way to create RegEx
			if len(got) != tc.lenOut || re.FindString(got) != "" {
				t.Errorf("got %s, want len %d", got, tc.lenOut)
			}
		})
	}
}
