package structs

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
