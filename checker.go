package checker

import (
	"github.com/AlkorMizar/Parentheses-cheker/structs"
)

const openBrace rune = ' '

func Check(input string) bool {
	bracesMap := map[rune]rune{
		rune('('): rune(openBrace),
		rune('['): rune(openBrace),
		rune('{'): rune(openBrace),
		rune(')'): rune('('),
		rune(']'): rune('['),
		rune('}'): rune('{'),
	}

	stack := structs.NewStack()

	for _, inpSymb := range input {
		if mapSymb, ok := bracesMap[inpSymb]; ok {
			if mapSymb == openBrace {
				stack.Push(inpSymb)
			} else if outSymb, err := stack.Pop(); err != nil || mapSymb != outSymb {
				return false
			}
		}
	}

	return true
}
