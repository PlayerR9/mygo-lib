package stack

import (
	"github.com/PlayerR9/mygo-lib/common"
)

// Capacity is a stack that has a capacity.
type Capacity[T any] struct {
	// stack is the underlying stack.
	stack Stack[T]

	// size is the number of elements in the stack.
	size uint

	// capacity is the maximum number of elements in the stack.
	capacity uint
}

// Push implements Stack.
func (c *Capacity[T]) Push(elem T) error {
	if c == nil {
		return common.ErrNilReceiver
	} else if c.size == c.capacity {
		return ErrFullStack
	}

	err := c.stack.Push(elem)
	if err != nil {
		return err
	}

	c.size++

	return nil
}

// Pop implements Stack.
func (c *Capacity[T]) Pop() (T, error) {
	if c == nil {
		return *new(T), common.ErrNilReceiver
	} else if c.size == 0 {
		return *new(T), ErrEmptyStack
	}

	top, _ := c.stack.Pop()
	c.size--

	return top, nil
}

// Peek implements Stack.
func (c Capacity[T]) Peek() (T, error) {
	if c.size == 0 {
		return *new(T), ErrEmptyStack
	}

	top, _ := c.stack.Peek()
	return top, nil
}

// IsEmpty implements Stack.
func (c Capacity[T]) IsEmpty() bool {
	return c.size == 0
}

// Size implements Stack.
func (c Capacity[T]) Size() uint {
	return c.size
}

// Free implements common.Type.
func (s *Capacity[T]) Free() {
	if s == nil {
		return
	}

	s.capacity = 0
	s.size = 0
	common.Free(s.stack)
	s.stack = nil
}

// WithCapacity creates a new stack with a specified capacity.
//
// Parameters:
//   - stack: The underlying stack to wrap with a capacity. If none is provided, a new
//     ArrayStack will be used.
//   - capacity: The maximum number of elements the stack can hold.
//
// Returns:
//   - Stack[T]: A new stack with the specified capacity.
//   - error: An error if the existing stack has more elements than the specified capacity.
func WithCapacity[T any](stack Stack[T], capacity uint) (*Capacity[T], error) {
	var size uint

	if stack == nil {
		stack = NewArrayStack[T]()
	} else {
		size = stack.Size()
		if size > capacity {
			return nil, ErrFullStack
		}
	}

	return &Capacity[T]{
		stack:    stack,
		size:     size,
		capacity: capacity,
	}, nil
}

// Reset resets the stack for reuse. Does nothing if the receiver is nil.
func (s *Capacity[T]) Reset() {
	if s == nil {
		return
	}

	stack, ok := s.stack.(interface{ Reset() })
	if ok {
		stack.Reset()
		s.size = 0

		return
	}

	for {
		_, err := s.Pop()
		if err != nil {
			break
		}
	}

	s.size = 0
}
