package structs

import "errors"

type Stack interface {
	Push(v rune)
	Pop() (rune, error)
}

type stack struct {
	s []rune
}

func NewStack() Stack {
	return &stack{make([]rune, 0)}
}

func (s *stack) Push(v rune) {
	s.s = append(s.s, v)
}

func (s *stack) Pop() (rune, error) {
	l := len(s.s)
	if l == 0 {
		return 0, errors.New("Empty Stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}
