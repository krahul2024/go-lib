package stack

import "errors"

type Stack[T comparable] struct {
	S         []T
	Sz        int
	ZeroValue T
}

func NewStack[T comparable](values ...T) *Stack[T] {
	return &Stack[T]{S: values, Sz: len(values)}
}

func (s *Stack[T]) Size() int {
	return s.Sz
}

func (s *Stack[T]) Push(v T) int {
	s.Sz += 1
	s.S = append(s.S, v)
	return s.Sz
}

func (s *Stack[T]) Top() (T, error) {
	if s.Sz == 0 {
		return s.ZeroValue, errors.New("empty stack")
	}
	return s.S[s.Sz-1], nil
}

func (s *Stack[T]) Pop() (T, error) {
	if s.Sz == 0 {
		return s.ZeroValue, errors.New("can't delete from empty stack")
	}
	value := s.S[s.Sz-1]
	s.S = s.S[:s.Sz-1]
	s.Sz -= 1
	return value, nil
}

func (s *Stack[T]) Empty() bool {
	if s.Sz == 0 {
		return true
	} else {
		return false
	}
}

func (s *Stack[T]) Clear() int {
	s.S = nil
	s.Sz = 0
	return s.Sz
}

func (s *Stack[T]) Has(value T) bool {
	for _, v := range s.S {
		if v == value {
			return true
		}
	}
	return false
}
