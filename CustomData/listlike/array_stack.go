package listlike

import (
	"slices"

	"github.com/PlayerR9/mygo-lib/common"
)

const (
	// NoCapacity is a sentinel value for a stack with no capacity.
	NoCapacity uint8 = 0xFF
)

// ArrayStack is a simple implementation of a stack that is backed by an array.
type ArrayStack struct {
	// slice is the internal slice.
	slice []any

	// capacity is the maximum capacity of the stack.
	capacity uint8
}

// Push implements the Stack interface.
func (s *ArrayStack) Push(elem any) error {
	if s == nil {
		return common.ErrNilReceiver
	} else if s.capacity != NoCapacity && uint8(len(s.slice)) == s.capacity {
		return ErrFullStack
	}

	s.slice = append(s.slice, elem)

	return nil
}

// Pop implements the Stack interface.
func (s *ArrayStack) Pop() (any, error) {
	if s == nil {
		return nil, common.ErrNilReceiver
	} else if len(s.slice) == 0 {
		return nil, ErrEmptyStack
	}

	elem := s.slice[len(s.slice)-1]
	s.slice = s.slice[:len(s.slice)-1]

	return elem, nil
}

// IsEmpty implements the Stack interface.
func (s ArrayStack) IsEmpty() bool {
	return len(s.slice) == 0
}

// NewArrayStack creates a new stack from a slice.
//
// Parameters:
//   - cap: The maximum capacity of the stack. If it is equal to NoCapacity, the stack
//     will have no capacity.
//   - elems: The elements to add to the stack.
//
// Returns:
//   - Stack: The new stack. Never returns nil.
//
// If the length of elems is larger than the given capacity, the capacity will be set
// to NoCapacity and the stack will have no capacity.
func NewArrayStack(cap uint8, elems ...any) Stack {
	stack := &ArrayStack{
		slice:    nil,
		capacity: cap,
	}

	_, err := stack.PushMany(elems)
	if err != nil {
		return nil
	}

	return stack
}

// PushMany adds multiple elements to the stack. If it has a capacity and the total
// length of the stack's underlying slice and the provided slice is larger than the
// capacity, the elements are truncated to fit the capacity.
//
// Parameters:
//   - elems: The elements to add.
//
// Returns:
//   - uint: The number of elements successfully pushed onto the stack.
//   - error: An error if the receiver is nil.
//
// Errors:
//   - common.ErrNilReceiver: If the receiver is nil.
func (s *ArrayStack) PushMany(elems []any) (uint, error) {
	lenElems := uint(len(elems))
	if lenElems == 0 {
		return 0, nil
	} else if s == nil {
		return 0, common.ErrNilReceiver
	}

	originalLen := uint(len(s.slice))
	totalLen := originalLen + lenElems

	var err error

	if s.capacity != NoCapacity && totalLen > uint(s.capacity) {
		lenElems = uint(s.capacity) - originalLen
		err = ErrFullStack
	}

	s.slice = append(s.slice, elems[:lenElems]...)
	slices.Reverse(s.slice[originalLen:])

	return lenElems, err
}

// Reset resets the stack for reuse. Does nothing if the receiver is nil.
func (s *ArrayStack) Reset() {
	if s == nil || len(s.slice) == 0 {
		return
	}

	clear(s.slice)

	if s.capacity == NoCapacity {
		s.slice = nil
	} else {
		s.slice = make([]any, 0, s.capacity)
	}
}

// Peek returns the element at the top of the stack without removing it.
//
// Returns:
//   - any: The element at the top of the stack. Nil if the stack is empty.
//   - error: An error of type ErrEmptyStack if the stack is empty.
func (s ArrayStack) Peek() (any, error) {
	if len(s.slice) == 0 {
		return nil, ErrEmptyStack
	}

	return s.slice[len(s.slice)-1], nil
}
