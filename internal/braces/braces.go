package braces

import (
	"math/rand"
	"strings"
)

type Braces struct{}

func NewBraces() *Braces {
	return &Braces{}
}

// business logic that generates strings with braces
func (b *Braces) Generate(leng int) string {
	var builder strings.Builder

	braces := []rune{'(', ')', '{', '}', '[', ']'}
	amount := len(braces)

	for leng > 0 {
		builder.WriteRune(braces[rand.Int31n(int32(amount))]) //nolint:gosec // this is used to simplify programm
		leng--
	}

	return builder.String()
}
