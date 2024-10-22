package listlike

import (
	"slices"

	"github.com/PlayerR9/mygo-lib/common"
)

// Stack is a simple implementation of a stack.
type Stack[T any] struct {
	// slice is the internal slice.
	slice []T
}

// Size implements the Lister interface.
func (s Stack[T]) Size() int {
	return len(s.slice)
}

// IsEmpty implements the Lister interface.
func (s Stack[T]) IsEmpty() bool {
	return len(s.slice) == 0
}

// Reset implements the Lister interface.
func (s *Stack[T]) Reset() {
	if s == nil {
		return
	}

	if len(s.slice) > 0 {
		zero := *new(T)

		for i := range s.slice {
			s.slice[i] = zero
		}

		s.slice = nil
	}
}

// NewStack creates a new stack from a slice. It is also possible
// use:
//
//	var stack Stack[T]
//
// To create an empty stack that is not a pointer.
//
// Parameters:
//   - elems: The elements to add to the stack.
//
// Returns:
//   - *Stack[T]: The new stack. Never returns nil.
//
// WARNING: As a side-effect, the original list will be reversed.
func NewStack[T any](elems []T) *Stack[T] {
	if len(elems) == 0 {
		return &Stack[T]{}
	}

	slices.Reverse(elems)

	return &Stack[T]{
		slice: elems,
	}
}

// Push adds an element to the stack.
//
// Parameters:
//   - elem: The element to add.
//
// Returns:
//   - error: An error if the receiver is nil.
func (s *Stack[T]) Push(elem T) error {
	if s == nil {
		return common.ErrNilReceiver
	}

	s.slice = append(s.slice, elem)

	return nil
}

// PushMany adds multiple elements to the stack. If it has at least
// one element but the receiver is nil, an error is returned.
//
// Parameters:
//   - elems: The elements to add.
//
// Returns:
//   - error: An error if the receiver is nil.
//
// WARNING: As a side-effect, the original list will be reversed.
func (s *Stack[T]) PushMany(elems []T) error {
	if len(elems) == 0 {
		return nil
	} else if s == nil {
		return common.ErrNilReceiver
	}

	slices.Reverse(elems)

	s.slice = append(s.slice, elems...)

	return nil
}

// Pop removes an element from the stack.
//
// Returns:
//   - T: The element that was removed.
//   - error: An error if the element could not be removed from the stack.
//
// Errors:
//   - ErrEmptyStack: If the stack is empty.
func (s *Stack[T]) Pop() (T, error) {
	if s == nil || len(s.slice) == 0 {
		return *new(T), ErrEmptyStack
	}

	elem := s.slice[len(s.slice)-1]
	s.slice = s.slice[:len(s.slice)-1]

	return elem, nil
}

// Peek returns the element at the top of the stack.
//
// Returns:
//   - T: The element at the top of the stack.
//   - error: An error if the stack is empty.
//
// Errors:
//   - ErrEmptyStack: If the stack is empty.
func (s Stack[T]) Peek() (T, error) {
	if len(s.slice) == 0 {
		return *new(T), ErrEmptyStack
	}

	return s.slice[len(s.slice)-1], nil
}
