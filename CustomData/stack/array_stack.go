package stack

import (
	"slices"

	"github.com/PlayerR9/mygo-lib/common"
)

// ArrayStack is a simple implementation of a stack that is backed by an array.
type ArrayStack[T any] struct {
	// slice is the backing array.
	slice []T

	// lenSlice is the number of elements in the slice.
	lenSlice uint
}

// Push implements Stack.
//
// Never returns ErrFullStack.
func (s *ArrayStack[T]) Push(elem T) error {
	if s == nil {
		return common.ErrNilReceiver
	}

	s.slice = append(s.slice, elem)
	s.lenSlice++

	return nil
}

// Pop implements Stack.
func (s *ArrayStack[T]) Pop() (T, error) {
	if s == nil {
		return *new(T), common.ErrNilReceiver
	}

	if s.lenSlice == 0 {
		return *new(T), ErrEmptyStack
	}

	top := s.slice[s.lenSlice-1]
	s.slice = s.slice[:s.lenSlice-1]
	s.lenSlice--

	return top, nil
}

// Peek implements Stack.
func (s ArrayStack[T]) Peek() (T, error) {
	if s.lenSlice == 0 {
		return *new(T), ErrEmptyStack
	}

	return s.slice[s.lenSlice-1], nil
}

// IsEmpty implements Stack.
func (s ArrayStack[T]) IsEmpty() bool {
	return s.lenSlice == 0
}

// Size implements Stack.
func (s ArrayStack[T]) Size() uint {
	return s.lenSlice
}

// Free implements common.Type.
func (s *ArrayStack[T]) Free() {
	if s == nil {
		return
	}

	if len(s.slice) > 0 {
		clear(s.slice)
		s.slice = nil
	}

	s.lenSlice = 0
}

// NewArrayStack creates a new stack from a slice.
//
// Parameters:
//   - elems: The elements to add to the stack.
//
// Returns:
//   - *ArrayStack[T]: The new stack. Never returns nil.
func NewArrayStack[T any](elems ...T) *ArrayStack[T] {
	if len(elems) == 0 {
		return &ArrayStack[T]{
			slice:    nil,
			lenSlice: 0,
		}
	}

	slice := make([]T, len(elems))
	copy(slice, elems)

	slices.Reverse(slice)

	return &ArrayStack[T]{
		slice:    slice,
		lenSlice: uint(len(elems)),
	}
}

// Reset resets the stack for reuse. Does nothing if the receiver is nil.
func (s *ArrayStack[T]) Reset() {
	if s == nil || len(s.slice) == 0 {
		return
	}

	clear(s.slice)
	s.slice = nil
	s.lenSlice = 0
}
