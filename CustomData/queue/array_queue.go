package queue

import (
	"sync"

	"github.com/PlayerR9/mygo-lib/common"
)

// ArrayQueue is a simple implementation of a queue that is backed by an array.
// This implementation is thread-safe.
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
func (s *ArrayQueue[T]) Enqueue(elem T) error {
	if s == nil {
		return common.ErrNilReceiver
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.slice = append(s.slice, elem)
	s.lenSlice++

	return nil
}

// Dequeue implements Queue.
func (s *ArrayQueue[T]) Dequeue() (T, error) {
	if s == nil {
		return *new(T), common.ErrNilReceiver
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.lenSlice == 0 {
		return *new(T), ErrEmptyQueue
	}

	top := s.slice[0]
	s.slice = s.slice[1:]
	s.lenSlice--

	return top, nil
}

// Front implements Queue.
func (s *ArrayQueue[T]) Front() (T, error) {
	if s == nil {
		return *new(T), common.ErrNilReceiver
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.lenSlice == 0 {
		return *new(T), ErrEmptyQueue
	}

	return s.slice[0], nil
}

// IsEmpty implements Queue.
func (s *ArrayQueue[T]) IsEmpty() bool {
	if s == nil {
		return true
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.lenSlice == 0
}

// Size implements Queue.
func (s *ArrayQueue[T]) Size() uint {
	if s == nil {
		return 0
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.lenSlice
}

// Free implements common.Type.
func (s *ArrayQueue[T]) Free() {
	if s == nil {
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	clear(s.slice)
	s.slice = nil
	s.lenSlice = 0
}

// NewArrayQueue creates a new queue from a slice.
//
// Parameters:
//   - elems: The elements to add to the queue.
//
// Returns:
//   - *ArrayQueue[T]: The new queue. Never returns nil.
func NewArrayQueue[T any](elems ...T) *ArrayQueue[T] {
	queue := new(ArrayQueue[T])
	if len(elems) == 0 {
		return queue
	}

	_, _ = queue.EnqueueMany(elems)
	return queue
}

// Reset resets the queue for reuse. Does nothing if the receiver is nil.
func (s *ArrayQueue[T]) Reset() {
	if s == nil {
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	clear(s.slice)
	s.slice = nil
	s.lenSlice = 0
}

// EnqueueMany adds multiple elements to the queue.
//
// Parameters:
//   - elems: A slice of elements to be added to the queue.
//
// Returns:
//   - uint: The number of elements successfully added to the queue.
//   - error: An error of type common.ErrNilReceiver if the receiver is nil.
func (s *ArrayQueue[T]) EnqueueMany(elems []T) (uint, error) {
	lenElems := uint(len(elems))
	if lenElems == 0 {
		return 0, nil
	} else if s == nil {
		return 0, common.ErrNilReceiver
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	slice := make([]T, len(elems))
	copy(slice, elems)

	s.slice = append(s.slice, slice...)
	s.lenSlice += lenElems

	return lenElems, nil
}
