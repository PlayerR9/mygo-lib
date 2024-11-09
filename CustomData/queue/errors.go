package queue

import "errors"

var (
	// ErrEmptyQueue occurs when a dequeue or front operation is called on an empty queue.
	// This error can be checked with the == operator.
	ErrEmptyQueue error

	// ErrFullQueue occurs when a enqueue operation is called on a full queue.
	// This error can be checked with the == operator.
	ErrFullQueue error
)

func init() {
	ErrEmptyQueue = errors.New("empty queue")
	ErrFullQueue = errors.New("full queue")
}
