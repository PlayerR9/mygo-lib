package queue

import (
	"sync"

	"github.com/PlayerR9/mygo-lib/common"
)

// ArrayQueue is a simple implementation of a queue that is backed by an array.
// This implementation is thread-safe.
//
// An empty array queue can be created using the `queue := new(ArrayQueue[T])` constructor.
type ArrayQueue[T any] struct {
	// slice is the backing array.
	slice []T

	// lenSlice is the number of elements in the slice.
	lenSlice uint

	// mu is the mutex.
	mu sync.RWMutex
}

// Enqueue implements Queue.
//
// Never returns ErrFullQueue.
func (q *ArrayQueue[T]) Enqueue(elem T) error {
	if q == nil {
		return common.ErrNilReceiver
	}

	q.mu.Lock()
	defer q.mu.Unlock()

	q.slice = append(q.slice, elem)
	q.lenSlice++

	return nil
}

// Dequeue implements Queue.
func (q *ArrayQueue[T]) Dequeue() (T, error) {
	if q == nil {
		return *new(T), common.ErrNilReceiver
	}

	q.mu.Lock()
	defer q.mu.Unlock()

	if q.lenSlice == 0 {
		return *new(T), ErrEmptyQueue
	}

	front := q.slice[0]
	q.slice = q.slice[1:]
	q.lenSlice--

	return front, nil
}

// Front implements Queue.
func (q *ArrayQueue[T]) Front() (T, error) {
	if q == nil {
		return *new(T), common.ErrNilReceiver
	}

	q.mu.RLock()
	defer q.mu.RUnlock()

	if q.lenSlice == 0 {
		return *new(T), ErrEmptyQueue
	}

	return q.slice[0], nil
}

// IsEmpty implements Queue.
func (q *ArrayQueue[T]) IsEmpty() bool {
	if q == nil {
		return true
	}

	q.mu.RLock()
	defer q.mu.RUnlock()

	return q.lenSlice == 0
}

// Size implements Queue.
func (q *ArrayQueue[T]) Size() uint {
	if q == nil {
		return 0
	}

	q.mu.RLock()
	defer q.mu.RUnlock()

	return q.lenSlice
}

// Free implements common.Typer.
func (q *ArrayQueue[T]) Free() {
	if q == nil {
		return
	}

	q.mu.Lock()
	defer q.mu.Unlock()

	clear(q.slice)
	q.slice = nil
	q.lenSlice = 0
}

// Reset implements common.Resetter.
func (q *ArrayQueue[T]) Reset() {
	if q == nil {
		return
	}

	q.mu.Lock()
	defer q.mu.Unlock()

	clear(q.slice)
	q.slice = nil
	q.lenSlice = 0
}

// EnqueueMany adds multiple elements to the queue in the order they are passed.
//
// Parameters:
//   - elems: A slice of elements to be added to the queue.
//
// Returns:
//   - uint: The number of elements successfully enqueued onto the queue.
//   - error: An error of type common.ErrNilReceiver if the receiver is nil.
func (q *ArrayQueue[T]) EnqueueMany(elems []T) (uint, error) {
	lenElems := uint(len(elems))
	if lenElems == 0 {
		return 0, nil
	} else if q == nil {
		return 0, common.ErrNilReceiver
	}

	q.mu.Lock()
	defer q.mu.Unlock()

	q.slice = append(q.slice, elems...)
	q.lenSlice += lenElems

	return lenElems, nil
}
