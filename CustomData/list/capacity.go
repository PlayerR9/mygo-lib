package list

import (
	"sync"

	"github.com/PlayerR9/mygo-lib/common"
)

// Capacity is a wrapper for a List that allows for a specified capacity.
type Capacity[T any] struct {
	// list is the underlying list.
	list List[T]

	// size is the number of elements in the list.
	size uint

	// capacity is the maximum number of elements in the list.
	capacity uint

	// mu is the mutex for the list.
	mu sync.RWMutex
}

// Enlist implements List.
//
// Errors:
//   - common.ErrInvalidObject: If the method `Free()` is called.
func (c *Capacity[T]) Enlist(elem T) error {
	if c == nil {
		return common.ErrNilReceiver
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.size == c.capacity {
		return ErrFullList
	} else if c.list == nil {
		// Method `Free()` was called.
		return common.NewErrInvalidObject("Enlist")
	}

	err := c.list.Enlist(elem)
	if err != nil {
		return err
	}

	c.size++

	return nil
}

// Prepend implements List.
//
// Errors:
//   - common.ErrInvalidObject: If the method `Free()` is called.
func (c *Capacity[T]) Prepend(elem T) error {
	if c == nil {
		return common.ErrNilReceiver
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.size == c.capacity {
		return ErrFullList
	} else if c.list == nil {
		// Method `Free()` was called.
		return common.NewErrInvalidObject("Prepend")
	}

	err := c.list.Prepend(elem)
	if err != nil {
		return err
	}

	c.size++

	return nil
}

// Delist implements List.
//
// Errors:
//   - common.ErrInvalidObject: If the method `Free()` is called.
func (c *Capacity[T]) Delist() (T, error) {
	if c == nil {
		return *new(T), common.ErrNilReceiver
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.size == 0 {
		return *new(T), ErrEmptyList
	} else if c.list == nil {
		// Method `Free()` was called.
		return *new(T), common.NewErrInvalidObject("Delist")
	}

	front, err := c.list.Delist()
	if err != nil {
		return *new(T), err
	}

	c.size--

	return front, nil
}

// Front implements List.
//
// Errors:
//   - common.ErrInvalidObject: If the method `Free()` is called.
func (c *Capacity[T]) Front() (T, error) {
	if c == nil {
		return *new(T), common.ErrNilReceiver
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.size == 0 {
		return *new(T), ErrEmptyList
	} else if c.list == nil {
		// Method `Free()` was called.
		return *new(T), common.NewErrInvalidObject("Front")
	}

	front, err := c.list.Front()
	return front, err
}

// Back implements List.
//
// Errors:
//   - common.ErrInvalidObject: If the method `Free()` is called.
func (c *Capacity[T]) Back() (T, error) {
	if c == nil {
		return *new(T), common.ErrNilReceiver
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.size == 0 {
		return *new(T), ErrEmptyList
	} else if c.list == nil {
		// Method `Free()` was called.
		return *new(T), common.NewErrInvalidObject("Back")
	}

	back, err := c.list.Back()
	return back, err
}

// IsEmpty implements List.
func (c *Capacity[T]) IsEmpty() bool {
	if c == nil {
		return true
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.size == 0
}

// Size implements List.
func (c *Capacity[T]) Size() uint {
	if c == nil {
		return 0
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.size
}

// Free implements List.
func (c *Capacity[T]) Free() {
	if c == nil {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	c.capacity = 0
	c.size = 0

	c.list.Free()
	c.list = nil
}

// Reset implements common.Resetter.
func (c *Capacity[T]) Reset() {
	if c == nil {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	Reset(c.list)

	c.size = 0
}

// WithCapacity creates a new list with a specified capacity.
//
// Parameters:
//   - list: The underlying list to wrap with a capacity.
//   - capacity: The maximum number of elements the list can hold.
//
// Returns:
//   - *Capacity[T]: A new list with the specified capacity.
//   - error: An error if the list could not be created.
//
// Errors:
//   - common.ErrBadParam: If the list parameter is nil.
//   - ErrFullList: If the existing list has more elements than the specified capacity.
func WithCapacity[T any](list List[T], capacity uint) (*Capacity[T], error) {
	if list == nil {
		return nil, common.NewErrNilParam("list")
	}

	size := list.Size()
	if size > capacity {
		return nil, ErrFullList
	}

	return &Capacity[T]{
		list:     list,
		size:     size,
		capacity: capacity,
	}, nil
}
