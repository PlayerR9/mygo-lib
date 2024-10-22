package listlike

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/PlayerR9/mygo-lib/common"
)

var (
	// ErrEmptyStack is an error that is returned when the stack is empty.
	//
	// Format:
	//
	//   "empty stack"
	ErrEmptyStack error
)

func init() {
	ErrEmptyStack = errors.New("empty stack")
}

// Stack is a stack.
//
// There is no need for a "new" constructor. Just use:
//
//	var stack Stack[T]
//
// To create an empty stack.
type Stack[T any] struct {
	// elems is the elements in the stack.
	elems []T
}

// Size implements the Lister interface.
func (s Stack[T]) Size() int {
	return len(s.elems)
}

// IsEmpty implements the Lister interface.
func (s Stack[T]) IsEmpty() bool {
	return len(s.elems) == 0
}

// IsNil implements the Lister interface.
func (s *Stack[T]) IsNil() bool {
	return s == nil
}

// String implements the fmt.Stringer interface.
func (s Stack[T]) String() string {
	elems := make([]string, 0, len(s.elems))

	for _, elem := range s.elems {
		elems = append(elems, fmt.Sprint(elem))
	}

	slices.Reverse(elems)

	return "Stack[ <-> " + strings.Join(elems, " <-> ") + "]"
}

// NewStack creates a new stack.
//
// Returns:
//   - *Stack[T]: The new stack. Never returns nil.
//
// This is just for convenience since it does the same as:
//
//	var stack Stack[T]
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// NewStackFromSlice creates a new stack from a slice.
//
// Parameters:
//   - elems: The elements to add to the stack.
//
// Returns:
//   - *Stack[T]: The new stack. Never returns nil.
func NewStackFromSlice[T any](slice []T) *Stack[T] {
	elems := make([]T, len(slice))
	copy(elems, slice)

	slices.Reverse(elems)

	return &Stack[T]{
		elems: elems,
	}
}

// Push adds an element to the stack.
//
// Parameters:
//   - elem: The element to add.
//
// Returns:
//   - error: An error if the element could not be added to the stack.
//
// Errors:
//   - errors.ErrNilReceiver: If the receiver is nil.
func (s *Stack[T]) Push(elem T) error {
	if s == nil {
		return common.ErrNilReceiver
	}

	s.elems = append(s.elems, elem)

	return nil
}

// PushMany adds multiple elements to the stack.
//
// Parameters:
//   - elems: The elements to add.
//
// Returns:
//   - error: An error if the element could not be added to the stack.
//
// Errors:
//   - errors.ErrNilReceiver: If the receiver is nil.
//
// WARNING: As a side-effect, the original list will be reversed.
func (s *Stack[T]) PushMany(elems []T) error {
	if len(elems) == 0 {
		return nil
	} else if s == nil {
		return common.ErrNilReceiver
	}

	slices.Reverse(elems)

	s.elems = append(s.elems, elems...)

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
	if s == nil || len(s.elems) == 0 {
		return *new(T), ErrEmptyStack
	}

	elem := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]

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
	if len(s.elems) == 0 {
		return *new(T), ErrEmptyStack
	}

	return s.elems[len(s.elems)-1], nil
}
