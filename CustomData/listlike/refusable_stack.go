package listlike

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

	lenSlice uint
}

// Size implements the Lister interface.
func (s RefusableStack[T]) Size() uint {
	return s.top
}

// IsEmpty implements the Lister interface.
func (s RefusableStack[T]) IsEmpty() bool {
	return s.top == 0
}

// Reset implements the Lister interface.
func (s *RefusableStack[T]) Reset() {
	if s == nil || s.lenSlice == 0 {
		return
	}

	clear(s.slice)
	s.slice = nil
	s.lenSlice = 0

	s.top = 0
}

// NewStack creates a new stack from a slice.
//
// Parameters:
//   - elems: The elements to add to the stack.
//
// Returns:
//   - *Stack[T]: The new stack. Never returns nil.
//
// WARNING: As a side-effect, the original list will be reversed.
func NewRefusableStack[T any](elems []T) *RefusableStack[T] {
	lenElems := uint(len(elems))
	if lenElems == 0 {
		return &RefusableStack[T]{
			slice:    nil,
			lenSlice: lenElems,
			top:      0,
		}
	}

	slices.Reverse(elems)

	return &RefusableStack[T]{
		slice:    elems,
		lenSlice: lenElems,
		top:      lenElems,
	}
}

// Push adds an element to the stack.
//
// Parameters:
//   - elem: The element to add.
//
// Returns:
//   - error: An error if the receiver is nil.
func (s *RefusableStack[T]) Push(elem T) error {
	if s == nil {
		return common.ErrNilReceiver
	}

	if s.top != s.lenSlice {
		return ErrCannotPush
	}

	s.slice = append(s.slice, elem)
	s.top++
	s.lenSlice++

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
func (s *RefusableStack[T]) PushMany(elems []T) error {
	if len(elems) == 0 {
		return nil
	} else if s == nil {
		return common.ErrNilReceiver
	}

	if s.top != s.lenSlice {
		return ErrCannotPush
	}

	slices.Reverse(elems)

	s.slice = append(s.slice, elems...)
	s.top += uint(len(elems))
	s.lenSlice += uint(len(elems))

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
func (s *RefusableStack[T]) Pop() (T, error) {
	if s == nil {
		return *new(T), ErrEmptyStack
	} else if s.top == 0 {
		return *new(T), ErrEmptyStack
	}

	s.top--

	return s.slice[s.top], nil
}

// Peek returns the element at the top of the stack.
//
// Returns:
//   - T: The element at the top of the stack.
//   - error: An error if the stack is empty.
//
// Errors:
//   - ErrEmptyStack: If the stack is empty.
func (s RefusableStack[T]) Peek() (T, error) {
	if s.top == 0 {
		return *new(T), ErrEmptyStack
	}

	return s.slice[s.top-1], nil
}

// Accept accepts all the elements that were popped. Does nothing if no element was popped.
func (s *RefusableStack[T]) Accept() {
	if s == nil {
		return
	}

	// common.Validate(s)

	if s.top != uint(len(s.slice)) {
		s.slice = s.slice[:s.top:s.top]
		s.lenSlice = s.top
	}
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
	if s == nil {
		return
	}

	if s.top != s.lenSlice {
		s.top++
	}
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
