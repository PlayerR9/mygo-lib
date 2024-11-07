package listlike

import (
	"github.com/PlayerR9/mygo-lib/common"
)

// Queue is an interface defining a queue.
type Queue interface {
	// IsEmpty checks whether the queue is empty.
	//
	// Returns:
	//   - bool: True if the queue is empty, false otherwise.
	IsEmpty() bool

	// Enqueue adds an element at the end of the queue.
	//
	// Parameters:
	//   - elem: The element to add.
	//
	// Returns:
	//   - error: An error of type common.ErrNilReceiver if the receiver is nil.
	Enqueue(elem any) error

	// Dequeue removes the first element from the queue.
	//
	// Returns:
	//   - any: The element that was removed. Nil if no element was removed.
	//   - error: An error if the element could not be removed from the queue.
	//
	// Errors:
	//   - ErrEmptyQueue: If the queue is empty.
	//   - common.ErrNilReceiver: If the receiver is nil.
	Dequeue() (any, error)

	// First returns the element at the start of the queue without removing it.
	//
	// Returns:
	//   - any: The element at the start of the queue. Nil if the queue is empty.
	//   - error: An error of type ErrEmptyQueue if the queue is empty.
	First() (any, error)
}

// ArrayQueue is a simple implementation of a queue. An empty queue can either be
// created with the `var queue ArrayQueue[T]` syntax or with the `new(ArrayQueue[T])`
type ArrayQueue[T any] struct {
	// slice is the internal slice.
	slice []T
}

// Enqueue implements the Queue interface.
func (q *ArrayQueue[T]) Enqueue(elem T) error {
	if q == nil {
		return common.ErrNilReceiver
	}

	q.slice = append(q.slice, elem)

	return nil
}

// IsEmpty implements the Queue interface.
func (q ArrayQueue[T]) IsEmpty() bool {
	return len(q.slice) == 0
}

// Dequeue implements the Queue interface.
func (q *ArrayQueue[T]) Dequeue() (T, error) {
	if q == nil {
		return *new(T), common.ErrNilReceiver
	} else if len(q.slice) == 0 {
		return *new(T), ErrEmptyQueue
	}

	elem := q.slice[0]
	q.slice = q.slice[1:]

	return elem, nil
}

// First implements the Queue interface.
func (q ArrayQueue[T]) First() (T, error) {
	if len(q.slice) == 0 {
		return *new(T), ErrEmptyQueue
	}

	return q.slice[0], nil
}

// NewArrayQueue creates a new queue from a slice.
//
// Parameters:
//   - elems: The elements to add to the queue.
//
// Returns:
//   - *Queue[T]: The new queue. Never returns nil.
func NewArrayQueue[T any](elems []T) *ArrayQueue[T] {
	if len(elems) == 0 {
		return &ArrayQueue[T]{}
	}

	return &ArrayQueue[T]{
		slice: elems,
	}
}

// EnqueueMany adds multiple elements to the queue. If it has at least
// one element but the receiver is nil, an error is returned.
//
// Parameters:
//   - elems: The elements to add.
//
// Returns:
//   - error: An error if the receiver is nil.
func (q *ArrayQueue[T]) EnqueueMany(elems []T) error {
	if len(elems) == 0 {
		return nil
	} else if q == nil {
		return common.ErrNilReceiver
	}

	q.slice = append(q.slice, elems...)

	return nil
}

// Size returns the number of elements in the queue.
//
// Returns:
//   - int: The number of elements in the queue. Never negative.
func (q ArrayQueue[T]) Size() uint {
	return uint(len(q.slice))
}

// Reset resets the queue for reuse. Does nothing if the receiver is nil.
func (q *ArrayQueue[T]) Reset() {
	if q == nil || len(q.slice) == 0 {
		return
	}

	clear(q.slice)
	q.slice = nil
}
