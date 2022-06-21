package checker

import "errors"

type Stack struct {
	s []rune
}

func NewStack() *Stack {
	return &Stack{make([]rune, 0)}
}

func (s *Stack) Push(v rune) {
	s.s = append(s.s, v)
}

func (s *Stack) Pop() (rune, error) {
	l := len(s.s)
	if l == 0 {
		return 0, errors.New("empty stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]

	return res, nil
}

const openBrace rune = ' '

// function that checks if given string is balnced
func Check(input string) bool {
	bracesMap := map[rune]rune{
		rune('('): openBrace,
		rune('['): openBrace,
		rune('{'): openBrace,
		rune(')'): rune('('),
		rune(']'): rune('['),
		rune('}'): rune('{'),
	}

	stack := NewStack()

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
