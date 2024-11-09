package queue

import (
	"slices"
	"sync"

	"github.com/PlayerR9/mygo-lib/common"
)

// Refuse is a wrapper for a Queue that allows elements to be refused, meaning that
// once dequeued we can decide whether to accept that dequeuing sequence or not. If not,
// they will be put back on the end of the queue.
type Refuse[T any] struct {
	// queue is the internal queue.
	queue Queue[T]

	// dequeued is the queue of elements that were dequeued.
	dequeued []T

	// mu is the mutex for the RefusableStack.
	mu sync.RWMutex
}

// Size implements the Queue interface.
//
// Panics:
//   - common.ErrInvalidObject: If the method `Free()` is called.
func (r *Refuse[T]) Size() uint {
	if r == nil {
		return 0
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.queue == nil {
		panic(common.NewErrInvalidObject("Size"))
	}

	size := r.queue.Size()
	return size
}

// IsEmpty implements the Queue interface.
//
// Panics:
//   - common.ErrInvalidObject: If the method `Free()` is called.
func (r *Refuse[T]) IsEmpty() bool {
	if r == nil {
		return true
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.queue == nil {
		panic(common.NewErrInvalidObject("IsEmpty"))
	}

	ok := r.queue.IsEmpty()
	return ok
}

// Enqueue implements the Queue interface.
//
// Errors:
//   - common.ErrInvalidObject: If the method `Free()` is called.
func (r *Refuse[T]) Enqueue(elem T) error {
	if r == nil {
		return common.ErrNilReceiver
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if r.queue == nil {
		return common.NewErrInvalidObject("Enqueue")
	}

	err := r.queue.Enqueue(elem)
	if err != nil {
		return err
	}

	return nil
}

// Dequeue implements the Queue interface.
//
// Errors:
//   - common.ErrInvalidObject: If the method `Free()` is called.
func (r *Refuse[T]) Dequeue() (T, error) {
	if r == nil {
		return *new(T), common.ErrNilReceiver
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if r.queue == nil {
		return *new(T), common.NewErrInvalidObject("Dequeue")
	}

	top, err := r.queue.Dequeue()
	if err != nil {
		return *new(T), err
	}

	r.dequeued = append(r.dequeued, top)

	return top, nil
}

// Front implements the Queue interface.
//
// Errors:
//   - common.ErrInvalidObject: If the method `Free()` is called.
func (r *Refuse[T]) Front() (T, error) {
	if r == nil {
		return *new(T), common.ErrNilReceiver
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.queue == nil {
		return *new(T), common.NewErrInvalidObject("Front")
	}

	top, err := r.queue.Front()
	if err != nil {
		return *new(T), err
	}

	return top, nil
}

// Reset implements common.Resetter.
func (r *Refuse[T]) Reset() {
	if r == nil {
		return
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if r.queue == nil {
		clear(r.dequeued)
		r.dequeued = nil

		return
	}

	Reset(r.queue)
	r.queue = nil

	clear(r.dequeued)
	r.dequeued = nil
}

// NewRefuse creates a new queue from a slice.
//
// Parameters:
//   - queue: The queue that can be refused.
//
// Returns:
//   - *Refuse[T]: The new queue. Nil if the queue is nil.
func NewRefuse[T any](queue Queue[T]) *Refuse[T] {
	if queue == nil {
		return nil
	}

	return &Refuse[T]{
		queue: queue,
	}
}

// Accept accepts all the elements that were dequeued. It's a no-op if no element was dequeued.
//
// Returns:
//   - error: An error of type common.ErrNilReceiver if the receiver is nil.
func (r *Refuse[T]) Accept() error {
	if r == nil {
		return common.ErrNilReceiver
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	clear(r.dequeued)
	r.dequeued = nil

	return nil
}

// Refuse refuses any element that was dequeued since the last time `Accept()` was called.
// It's a no-op if no element was dequeued.
//
// Returns:
//   - error: An error if refusing the element failed.
//
// Errors:
//   - common.ErrNilReceiver: If the receiver is nil.
//   - common.ErrInvalidObject: If the method `Free()` is called.
func (r *Refuse[T]) Refuse() error {
	if r == nil {
		return common.ErrNilReceiver
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if r.queue == nil {
		return common.NewErrInvalidObject("Refuse")
	}

	var i int
	var err error

	for i = 0; i < len(r.dequeued) && err == nil; i++ {
		err = r.queue.Enqueue(r.dequeued[i])
	}

	if i == len(r.dequeued) {
		clear(r.dequeued)
		r.dequeued = nil
	} else {
		clear(r.dequeued[:i])
		r.dequeued = r.dequeued[i:]
	}

	return err
}

// RefuseOne refuses the last dequeued element. It's a no-op if no element was dequeued.
//
// Returns:
//   - error: An error if refusing the element failed.
//
// Errors:
//   - common.ErrNilReceiver: If the receiver is nil.
//   - common.ErrInvalidObject: If the method `Free()` is called.
func (r *Refuse[T]) RefuseOne() error {
	if r == nil {
		return common.ErrNilReceiver
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if len(r.dequeued) == 0 {
		return nil
	} else if r.queue == nil {
		return common.NewErrInvalidObject("RefuseOne")
	}

	top := r.dequeued[0]
	r.dequeued = r.dequeued[1:]

	err := r.queue.Enqueue(top)
	return err
}

// Dequeued returns the elements that were dequeued from the queue since the last
// `Accept()` or `Refuse()` operation. The returned slice contains the elements
// in the order they were dequeued, with the most recently dequeued element at the
// first position.
//
// Returns:
//   - []T: The elements that were dequeued.
func (r *Refuse[T]) Dequeued() []T {
	if r == nil {
		return nil
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	if len(r.dequeued) == 0 {
		return nil
	}

	slice := make([]T, len(r.dequeued))
	copy(slice, r.dequeued)

	slices.Reverse(slice)

	return slice
}
