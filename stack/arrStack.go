package stack

import "errors"

type ArrStack[T comparable] struct {
	S         []T
	Sz        int
	ZeroValue T
}

func NewArrStack[T comparable](values ...T) *ArrStack[T] {
	return &ArrStack[T]{S: values, Sz: len(values)}
}

func (s *ArrStack[T]) Size() int {
	return s.Sz
}

func (s *ArrStack[T]) Push(v T) int {
	s.Sz += 1
	s.S = append(s.S, v)
	return s.Sz
}

func (s *ArrStack[T]) Top() (T, error) {
	if s.Sz == 0 {
		return s.ZeroValue, errors.New("empty ArrStack")
	}
	return s.S[s.Sz-1], nil
}

func (s *ArrStack[T]) Pop() (T, error) {
	if s.Sz == 0 {
		return s.ZeroValue, errors.New("can't delete from empty ArrStack")
	}
	value := s.S[s.Sz-1]
	s.S = s.S[:s.Sz-1]
	s.Sz -= 1
	return value, nil
}

func (s *ArrStack[T]) Empty() bool {
	if s.Sz == 0 {
		return true
	} else {
		return false
	}
}

func (s *ArrStack[T]) Clear() int {
	s.S = nil
	s.Sz = 0
	return s.Sz
}

func (s *ArrStack[T]) Has(value T) bool {
	for _, v := range s.S {
		if v == value {
			return true
		}
	}
	return false
}
