package checker

import (
	"github.com/AlkorMizar/Parentheses-cheker/structs"
)

type Stack interface {
	Push(v rune)
	Pop() (rune, error)
}

const openBrace rune = ' '

func Check(input string) bool {
	bracesMap := map[rune]rune{
		rune('('): openBrace,
		rune('['): openBrace,
		rune('{'): openBrace,
		rune(')'): rune('('),
		rune(']'): rune('['),
		rune('}'): rune('{'),
	}

	var stack Stack = structs.NewStack()

	for _, inpSymb := range input {
		if mapSymb, ok := bracesMap[inpSymb]; ok {
			if mapSymb == openBrace {
				stack.Push(inpSymb)
			} else if outSymb, err := stack.Pop(); err != nil || mapSymb != outSymb {
				return false
			}
		}
	}

	_, err := stack.Pop()

	return err != nil
}
