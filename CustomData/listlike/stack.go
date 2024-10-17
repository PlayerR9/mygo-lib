package listlike

import (
	"errors"
	"slices"

	gers "github.com/PlayerR9/mygo-lib/errors"
)

var (
	ErrEmptyStack error
)

func init() {
	ErrEmptyStack = errors.New("empty stack")
}

type Stack[T any] struct {
	elems []T
}

func (s *Stack[T]) Push(elem T) error {
	if s == nil {
		return gers.ErrNilReceiver
	}

	s.elems = append(s.elems, elem)

	return nil
}

func (s *Stack[T]) PushMany(elems []T) error {
	if len(elems) == 0 {
		return nil
	} else if s == nil {
		return gers.ErrNilReceiver
	}

	slices.Reverse(elems)

	s.elems = append(s.elems, elems...)

	return nil
}

func (s *Stack[T]) Pop() (T, error) {
	if s == nil || len(s.elems) == 0 {
		return *new(T), ErrEmptyStack
	}

	elem := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]

	return elem, nil
}

func (s Stack[T]) Peek() (T, error) {
	if len(s.elems) == 0 {
		return *new(T), ErrEmptyStack
	}

	return s.elems[len(s.elems)-1], nil
}
