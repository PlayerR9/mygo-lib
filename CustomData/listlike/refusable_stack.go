package listlike

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	assert "github.com/PlayerR9/go-verify"
	gers "github.com/PlayerR9/mygo-lib/errors"
)

var (
	// ErrNotAccepted is an error that is returned when the stack is not accepted (or refused)
	// while a push operation is attempted.
	//
	// Format:
	//
	//   "stack was not accepted nor refused"
	ErrNotAccepted error
)

func init() {
	ErrNotAccepted = errors.New("stack was not accepted nor refused")
}

// RefusableStack is a stack that can be refused and accepted. This means we can pop elements
// without popping them from the stack.
type RefusableStack[T any] struct {
	// elems is the list of elements in the stack.
	elems []T

	// stack_size is the number of elements in the stack.
	stack_size int

	// stack_end is the index of the end of the stack.
	stack_end int
}

// Size implements the Lister interface.
func (rs RefusableStack[T]) Size() int {
	err := rs.Validate()
	assert.Err(err, "rs.Validate()")

	return rs.stack_size
}

// IsEmpty implements the Lister interface.
func (rs RefusableStack[T]) IsEmpty() bool {
	return rs.stack_size == 0
}

// IsNil implements the Lister interface.
func (rs *RefusableStack[T]) IsNil() bool {
	return rs == nil
}

// Validate implements the assert.Validater interface.
func (rs RefusableStack[T]) Validate() error {
	if rs.stack_size != len(rs.elems) {
		return fmt.Errorf("%q mismatch", "stack_size")
	}

	if rs.stack_end < 0 || rs.stack_end > rs.stack_size {
		return fmt.Errorf("%q is not in [%d, %d]", "stack_end", 0, rs.stack_size)
	}

	return nil
}

// String implements the fmt.Stringer interface.
func (rs RefusableStack[T]) String() string {
	elems := make([]string, 0, len(rs.elems))

	for _, elem := range rs.elems {
		elems = append(elems, fmt.Sprint(elem))
	}

	slices.Reverse(elems)

	return "RefusableStack[ <-> " + strings.Join(elems, " <-> ") + "]"
}

// NewRefusableStack creates a new RefusableStack.
//
// Returns:
//   - *RefusableStack[T]: The new RefusableStack. Never returns nil.
//
// This is just for convenience since it does the same as:
//
//	var stack RefusableStack[T]
func NewRefusableStack[T any]() *RefusableStack[T] {
	return &RefusableStack[T]{}
}

// NewRefusableStackFromSlice creates a new RefusableStack from a slice.
//
// Returns:
//   - *RefusableStack[T]: The new RefusableStack. Never returns nil.
func NewRefusableStackFromSlice[T any](slice []T) *RefusableStack[T] {
	elems := make([]T, len(slice))
	copy(elems, slice)

	slices.Reverse(elems)

	return &RefusableStack[T]{
		elems:      elems,
		stack_size: len(elems),
		stack_end:  len(elems),
	}
}

// Peek returns the element at the top of the stack.
//
// Returns:
//   - T: The element at the top of the stack.
//
// Errors:
//   - ErrEmptyStack: If the stack is empty.
func (rs RefusableStack[T]) Peek() (T, error) {
	err := rs.Validate()
	assert.Err(err, "rs.Validate()")

	if rs.stack_end == 0 {
		return *new(T), ErrEmptyStack
	}

	return rs.elems[rs.stack_end-1], nil
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
//   - ErrNotAccepted: If the stack was not accepted (or refused).
func (rs *RefusableStack[T]) Push(elem T) error {
	if rs == nil {
		return gers.ErrNilReceiver
	}

	err := rs.Validate()
	assert.Err(err, "rs.Validate()")

	if rs.stack_end != rs.stack_size {
		return ErrNotAccepted
	}

	rs.elems = append(rs.elems, elem)
	rs.stack_size++
	rs.stack_end++

	return nil
}

// Pop removes an element from the stack.
//
// Returns:
//   - T: The element that was removed.
//   - error: An error if the element could not be removed from the stack.
//
// Errors:
//   - errors.ErrNilReceiver: If the receiver is nil.
//   - ErrEmptyStack: If the stack is empty.
func (rs *RefusableStack[T]) Pop() (T, error) {
	zero := *new(T)

	if rs == nil {
		return zero, ErrEmptyStack
	}

	err := rs.Validate()
	assert.Err(err, "rs.Validate()")

	if rs.stack_end == 0 {
		return zero, ErrEmptyStack
	}

	rs.stack_end--

	return rs.elems[rs.stack_end], nil
}

// Accept accepts the stack. Does nothing if no element was popped.
func (rs *RefusableStack[T]) Accept() {
	if rs == nil {
		return
	}

	err := rs.Validate()
	assert.Err(err, "rs.Validate()")

	rs.elems = rs.elems[:rs.stack_end]
	rs.stack_size = rs.stack_end
}

// Refuse refuses the stack. Does nothing if no element was popped.
func (rs *RefusableStack[T]) Refuse() {
	if rs == nil {
		return
	}

	err := rs.Validate()
	assert.Err(err, "rs.Validate()")

	rs.stack_end = rs.stack_size
}

// Popped returns the elements that were popped from the stack.
//
// Returns:
//   - []T: The elements that were popped. Nil if no elements were popped.
//
// The first element in the returned slice is the last element that was popped.
func (rs RefusableStack[T]) Popped() []T {
	err := rs.Validate()
	assert.Err(err, "rs.Validate()")

	diff := rs.stack_size - rs.stack_end

	if diff == 0 {
		return nil
	}

	slice := make([]T, 0, diff)

	for _, elem := range rs.elems[rs.stack_end:] {
		slice = append(slice, elem)
	}

	return slice
}
