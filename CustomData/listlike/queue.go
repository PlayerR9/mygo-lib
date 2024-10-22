package listlike

import (
	"github.com/PlayerR9/mygo-lib/common"
)

// Queue is a simple implementation of a queue.
type Queue[T any] struct {
	// slice is the internal slice.
	slice []T
}

// Size implements the Lister interface.
func (s Queue[T]) Size() int {
	return len(s.slice)
}

// IsEmpty implements the Lister interface.
func (s Queue[T]) IsEmpty() bool {
	return len(s.slice) == 0
}

// Reset implements the Lister interface.
func (s *Queue[T]) Reset() {
	if s == nil {
		return
	}

	if len(s.slice) > 0 {
		zero := *new(T)

		for i := range s.slice {
			s.slice[i] = zero
		}

		s.slice = nil
	}
}

// NewQueue creates a new queue from a slice. It is also possible
// use:
//
//	var queue Queue[T]
//
// To create an empty queue that is not a pointer.
//
// Parameters:
//   - elems: The elements to add to the queue.
//
// Returns:
//   - *Queue[T]: The new queue. Never returns nil.
func NewQueue[T any](elems []T) *Queue[T] {
	if len(elems) == 0 {
		return &Queue[T]{}
	}

	return &Queue[T]{
		slice: elems,
	}
}

// Enqueue adds an element to the queue.
//
// Parameters:
//   - elem: The element to add.
//
// Returns:
//   - error: An error if the receiver is nil.
func (s *Queue[T]) Enqueue(elem T) error {
	if s == nil {
		return common.ErrNilReceiver
	}

	s.slice = append(s.slice, elem)

	return nil
}

// EnqueueMany adds multiple elements to the queue. If it has at least
// one element but the receiver is nil, an error is returned.
//
// Parameters:
//   - elems: The elements to add.
//
// Returns:
//   - error: An error if the receiver is nil.
func (s *Queue[T]) EnqueueMany(elems []T) error {
	if len(elems) == 0 {
		return nil
	} else if s == nil {
		return common.ErrNilReceiver
	}

	s.slice = append(s.slice, elems...)

	return nil
}

// Dequeue removes the first element from the queue.
//
// Returns:
//   - T: The element that was removed.
//   - error: An error if the element could not be removed from the queue.
//
// Errors:
//   - ErrEmptyQueue: If the queue is empty.
func (s *Queue[T]) Dequeue() (T, error) {
	if s == nil || len(s.slice) == 0 {
		return *new(T), ErrEmptyQueue
	}

	elem := s.slice[0]
	s.slice = s.slice[1:]

	return elem, nil
}

// First returns the element at the start of the queue.
//
// Returns:
//   - T: The element at the start of the queue.
//   - error: An error if the queue is empty.
//
// Errors:
//   - ErrEmptyQueue: If the queue is empty.
func (s Queue[T]) First() (T, error) {
	if len(s.slice) == 0 {
		return *new(T), ErrEmptyQueue
	}

	return s.slice[len(s.slice)-1], nil
}
