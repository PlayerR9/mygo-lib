package listlike

import (
	"errors"
	"fmt"
	"strings"

	"github.com/PlayerR9/mygo-lib/common"
)

var (
	// ErrEmptyQueue is an error that is returned when the queue is empty.
	//
	// Format:
	//
	//   "empty queue"
	ErrEmptyQueue error
)

func init() {
	ErrEmptyQueue = errors.New("empty queue")
}

// Queue is a queue.
type Queue[T any] struct {
	// elems is the list of elements in the queue.
	elems []T
}

// Size implements the Lister interface.
func (q Queue[T]) Size() int {
	return len(q.elems)
}

// IsEmpty implements the Lister interface.
func (q Queue[T]) IsEmpty() bool {
	return len(q.elems) == 0
}

// IsNil implements the Lister interface.
func (q *Queue[T]) IsNil() bool {
	return q == nil
}

// String implements the fmt.Stringer interface.
func (q Queue[T]) String() string {
	elems := make([]string, 0, len(q.elems))

	for _, elem := range q.elems {
		elems = append(elems, fmt.Sprint(elem))
	}

	return "Queue[ <- " + strings.Join(elems, " <- ") + " <- ]"
}

// NewQueue returns a new queue.
//
// Returns:
//   - *Queue[T]: The new queue. Never returns nil.
//
// This is just for convenience since it does the same as:
//
//	var queue Queue[T]
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

// NewQueueWithValues returns a new queue with the given values.
//
// Returns:
//   - *Queue[T]: The new queue. Never returns nil.
func NewQueueWithValues[T any](elems []T) *Queue[T] {
	return &Queue[T]{
		elems: elems,
	}
}

// Enqueue adds an element to the queue.
//
// Parameters:
//   - elem: The element to add.
//
// Returns:
//   - error: An error if the receiver is nil.
func (q *Queue[T]) Enqueue(elem T) error {
	if q == nil {
		return common.ErrNilReceiver
	}

	q.elems = append(q.elems, elem)

	return nil
}

// EnqueueMany adds multiple elements to the queue.
//
// Parameters:
//   - elems: The elements to add.
//
// Returns:
//   - error: An error if the receiver is nil.
func (q *Queue[T]) EnqueueMany(elems []T) error {
	if len(elems) == 0 {
		return nil
	} else if q == nil {
		return common.ErrNilReceiver
	}

	q.elems = append(q.elems, elems...)

	return nil
}

// Dequeue removes an element from the queue.
//
// Returns:
//   - T: The element that was removed.
//   - error: An error if the receiver is nil or the queue is empty.
func (q *Queue[T]) Dequeue() (T, error) {
	if q == nil || len(q.elems) == 0 {
		return *new(T), ErrEmptyQueue
	}

	elem := q.elems[0]
	q.elems = q.elems[1:]

	return elem, nil
}
