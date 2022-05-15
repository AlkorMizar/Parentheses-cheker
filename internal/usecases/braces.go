package usecases

import (
	"math/rand"
	"strings"
)

type braces struct{}

func NewBraces() *braces {
	return &braces{}
}

// business logic that generates strings with braces
func (b *braces) Generate(leng int) string {
	var builder strings.Builder

	braces := []rune{'(', ')', '{', '}', '[', ']'}
	amount := len(braces)

	for leng > 0 {
		builder.WriteRune(braces[rand.Int31n(int32(amount))]) //nolint:gosec // this is used to simplify programm
		leng--
	}

	return builder.String()
}
