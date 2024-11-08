package stack

import (
	"slices"

	"github.com/PlayerR9/mygo-lib/common"
)

// RefusableStack is a simple implementation of a stack. An empty stack can either be
// created with the `var stack Stack[T]` syntax or with the `new(Stack[T])`
// constructor.
type RefusableStack[T any] struct {
	// slice is the internal slice of the stack.
	slice []T

	// top is the top of the stack.
	top uint

	// lenSlice is the length of the slice.
	lenSlice uint
}

// Size implements the Stack interface.
func (s RefusableStack[T]) Size() uint {
	return s.top
}

// IsEmpty implements the Stack interface.
func (s RefusableStack[T]) IsEmpty() bool {
	return s.top == 0
}

// Push implements the Stack interface.
//
// Errors:
//   - ErrCannotPush: If the stack is not accepted nor refused yet.
func (s *RefusableStack[T]) Push(elem T) error {
	if s == nil {
		return common.ErrNilReceiver
	} else if s.top != s.lenSlice {
		return ErrCannotPush
	}

	s.slice = append(s.slice, elem)
	s.top++
	s.lenSlice++

	return nil
}

// Pop implements the Stack interface.
func (s *RefusableStack[T]) Pop() (T, error) {
	if s == nil {
		return *new(T), common.ErrNilReceiver
	} else if s.top == 0 {
		return *new(T), ErrEmptyStack
	}

	s.top--

	return s.slice[s.top], nil
}

// Peek implements the Stack interface.
func (s RefusableStack[T]) Peek() (T, error) {
	if s.top == 0 {
		return *new(T), ErrEmptyStack
	}

	return s.slice[s.top-1], nil
}

// NewStack creates a new stack from a slice.
//
// Parameters:
//   - elems: The elements to add to the stack.
//
// Returns:
//   - *RefusableStack[T]: The new stack. Never returns nil.
func NewRefusableStack[T any](elems ...T) *RefusableStack[T] {
	lenElems := uint(len(elems))
	if lenElems == 0 {
		return &RefusableStack[T]{
			slice:    nil,
			lenSlice: 0,
			top:      0,
		}
	}

	slice := make([]T, lenElems)
	copy(slice, elems)
	slices.Reverse(slice)

	return &RefusableStack[T]{
		slice:    slice,
		lenSlice: lenElems,
		top:      lenElems,
	}
}

// Accept accepts all the elements that were popped. Does nothing if no element was popped.
func (s *RefusableStack[T]) Accept() {
	if s == nil || s.top == s.lenSlice {
		return
	}

	s.slice = s.slice[:s.top:s.top]
	s.lenSlice = s.top
}

// Refuse refuses any element that was popped since the last time Accept was called.
// Does nothing if no element was popped.
func (s *RefusableStack[T]) Refuse() {
	if s == nil {
		return
	}

	s.top = s.lenSlice
}

// RefuseOne refuses the last popped element. Does nothing if no element was popped.
func (s *RefusableStack[T]) RefuseOne() {
	if s == nil || s.top == s.lenSlice {
		return
	}

	s.top++
}

// Popped returns the elements that were popped from the stack since the last
// Accept or Refuse operation. The returned slice contains the elements in the
// order they were popped, with the most recently popped element at the first
// position.
//
// Returns:
//   - []T: The elements that were popped. Nil if no elements were popped.
func (s RefusableStack[T]) Popped() []T {
	if s.top == s.lenSlice {
		return nil
	}

	slice := make([]T, s.lenSlice-s.top)
	copy(slice, s.slice[s.top:])

	return slice
}

// Reset resets the stack for reuse. Does nothing if the receiver is nil.
func (s *RefusableStack[T]) Reset() {
	if s == nil || s.lenSlice == 0 {
		return
	}

	clear(s.slice)
	s.slice = nil
	s.lenSlice = 0

	s.top = 0
}
