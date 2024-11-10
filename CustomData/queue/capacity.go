package queue

import (
	"sync"

	"github.com/PlayerR9/mygo-lib/common"
)

// Capacity is a wrapper for a Queue that allows for a specified capacity.
type Capacity[T any] struct {
	// queue is the underlying queue.
	queue Queue[T]

	// size is the number of elements in the queue.
	size uint

	// capacity is the maximum number of elements in the queue.
	capacity uint

	// mu is the mutex for the queue.
	mu sync.RWMutex
}

// Enqueue implements Queue.
//
// Errors:
//   - common.ErrInvalidObject: If the method `Free()` is called.
func (c *Capacity[T]) Enqueue(elem T) error {
	if c == nil {
		return common.ErrNilReceiver
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.size == c.capacity {
		return ErrFullQueue
	} else if c.queue == nil {
		// Method `Free()` was called.
		return common.NewErrInvalidObject("Enqueue")
	}

	err := c.queue.Enqueue(elem)
	if err != nil {
		return err
	}

	c.size++

	return nil
}

// Dequeue implements Queue.
//
// Errors:
//   - common.ErrInvalidObject: If the method `Free()` is called.
func (c *Capacity[T]) Dequeue() (T, error) {
	if c == nil {
		return *new(T), common.ErrNilReceiver
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.size == 0 {
		return *new(T), ErrEmptyQueue
	} else if c.queue == nil {
		// Method `Free()` was called.
		return *new(T), common.NewErrInvalidObject("Dequeue")
	}

	front, err := c.queue.Dequeue()
	if err != nil {
		return *new(T), err
	}

	c.size--

	return front, nil
}

// Front implements Queue.
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
		return *new(T), ErrEmptyQueue
	} else if c.queue == nil {
		// Method `Free()` was called.
		return *new(T), common.NewErrInvalidObject("Front")
	}

	front, err := c.queue.Front()
	return front, err
}

// IsEmpty implements Queue.
func (c *Capacity[T]) IsEmpty() bool {
	if c == nil {
		return true
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.size == 0
}

// Size implements Queue.
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

	Free(c.queue)

	c.queue = nil
}

// Reset implements common.Resetter.
func (c *Capacity[T]) Reset() {
	if c == nil {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	Reset(c.queue)

	c.size = 0
}

// WithCapacity creates a new queue with a specified capacity.
//
// Parameters:
//   - queue: The underlying queue to wrap with a capacity.
//   - capacity: The maximum number of elements the queue can hold.
//
// Returns:
//   - *Capacity[T]: A new queue with the specified capacity.
//   - error: An error if the queue could not be created.
//
// Errors:
//   - common.ErrBadParam: If the queue parameter is nil.
//   - ErrFullQueue: If the existing queue has more elements than the specified capacity.
func WithCapacity[T any](queue Queue[T], capacity uint) (*Capacity[T], error) {
	if queue == nil {
		return nil, common.NewErrNilParam("queue")
	}

	size := queue.Size()
	if size > capacity {
		return nil, ErrFullQueue
	}

	return &Capacity[T]{
		queue:    queue,
		size:     size,
		capacity: capacity,
	}, nil
}
