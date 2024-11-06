package listlike

import (
	"slices"

	"github.com/PlayerR9/mygo-lib/common"
)

const (
	// NoCapacity is a sentinel value for a stack with no capacity.
	NoCapacity uint8 = 0xFF
)

// ArrayStack is a simple implementation of a stack that is backed by an array. An empty
// stack can either be created with the `var stack ArrayStack` syntax or with the
// `new(ArrayStack)` constructor.
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

// Peek implements the Stack interface.
func (s ArrayStack) Peek() (any, error) {
	if len(s.slice) == 0 {
		return nil, ErrEmptyStack
	}

	return s.slice[len(s.slice)-1], nil
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
	if len(elems) == 0 {
		return &ArrayStack{
			slice:    nil,
			capacity: cap,
		}
	}

	if cap != NoCapacity && len(elems) >= int(NoCapacity) {
		cap = NoCapacity
	}

	var slice []any

	if cap == NoCapacity {
		slice = make([]any, len(elems))
	} else {
		slice = make([]any, len(elems), cap)
	}

	copy(slice, elems)
	slices.Reverse(slice[:len(elems)])

	return &ArrayStack{
		slice:    slice,
		capacity: cap,
	}
}

func (s *ArrayStack) PushMany(elems []any) (uint, error) {
	lenElems := uint(len(elems))
	if lenElems == 0 {
		return 0, nil
	} else if s == nil {
		return 0, common.ErrNilReceiver
	}

	for i := uint(0); i < lenElems; i++ {
		if s.capacity != NoCapacity && uint8(len(s.slice)) == s.capacity {
			return i, ErrFullStack
		}

		s.slice = append(s.slice, elems[i])
	}

	return lenElems, nil
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
