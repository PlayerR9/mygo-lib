package stack

import (
	"sync"

	"github.com/PlayerR9/mygo-lib/common"
)

// Capacity is a wrapper for a Stack that allows for a specified capacity.
type Capacity[T any] struct {
	// stack is the underlying stack.
	stack Stack[T]

	// size is the number of elements in the stack.
	size uint

	// capacity is the maximum number of elements in the stack.
	capacity uint

	// mu is the mutex for the stack.
	mu sync.RWMutex
}

// Push implements Stack.
//
// Errors:
//   - common.ErrInvalidObject: If the method `Free()` is called.
func (c *Capacity[T]) Push(elem T) error {
	if c == nil {
		return common.ErrNilReceiver
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.size == c.capacity {
		return ErrFullStack
	} else if c.stack == nil {
		// Method `Free()` was called.
		return common.NewErrInvalidObject("Push")
	}

	err := c.stack.Push(elem)
	if err != nil {
		return err
	}

	c.size++

	return nil
}

// Pop implements Stack.
//
// Errors:
//   - common.ErrInvalidObject: If the method `Free()` is called.
func (c *Capacity[T]) Pop() (T, error) {
	if c == nil {
		return *new(T), common.ErrNilReceiver
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.size == 0 {
		return *new(T), ErrEmptyStack
	} else if c.stack == nil {
		// Method `Free()` was called.
		return *new(T), common.NewErrInvalidObject("Pop")
	}

	top, err := c.stack.Pop()
	if err != nil {
		return *new(T), err
	}

	c.size--

	return top, nil
}

// Peek implements Stack.
//
// Errors:
//   - common.ErrInvalidObject: If the method `Free()` is called.
func (c *Capacity[T]) Peek() (T, error) {
	if c == nil {
		return *new(T), common.ErrNilReceiver
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.size == 0 {
		return *new(T), ErrEmptyStack
	} else if c.stack == nil {
		// Method `Free()` was called.
		return *new(T), common.NewErrInvalidObject("Peek")
	}

	top, err := c.stack.Peek()
	return top, err
}

// IsEmpty implements Stack.
func (c *Capacity[T]) IsEmpty() bool {
	if c == nil {
		return true
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.size == 0
}

// Size implements Stack.
func (c *Capacity[T]) Size() uint {
	if c == nil {
		return 0
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.size
}

// Free implements common.Type.
func (c *Capacity[T]) Free() {
	if c == nil {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	c.capacity = 0
	c.size = 0

	Free(c.stack)

	c.stack = nil
}

// Reset implements common.Resetter.
func (c *Capacity[T]) Reset() {
	if c == nil {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	Reset(c.stack)

	c.size = 0
}

// WithCapacity creates a new stack with a specified capacity.
//
// Parameters:
//   - stack: The underlying stack to wrap with a capacity.
//   - capacity: The maximum number of elements the stack can hold.
//
// Returns:
//   - *Capacity[T]: A new stack with the specified capacity.
//   - error: An error if the stack could not be created.
//
// Errors:
//   - common.ErrBadParam: If the stack parameter is nil.
//   - ErrFullStack: If the existing stack has more elements than the specified capacity.
func WithCapacity[T any](stack Stack[T], capacity uint) (*Capacity[T], error) {
	if stack == nil {
		return nil, common.NewErrNilParam("stack")
	}

	size := stack.Size()
	if size > capacity {
		return nil, ErrFullStack
	}

	return &Capacity[T]{
		stack:    stack,
		size:     size,
		capacity: capacity,
	}, nil
}
