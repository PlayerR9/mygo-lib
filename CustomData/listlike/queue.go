package listlike

import (
	"errors"

	gers "github.com/PlayerR9/mygo-lib/errors"
)

var (
	ErrEmptyQueue error
)

func init() {
	ErrEmptyQueue = errors.New("empty queue")
}

type Queue[T any] struct {
	elems []T
}

func (q *Queue[T]) Enqueue(elem T) error {
	if q == nil {
		return gers.ErrNilReceiver
	}

	q.elems = append(q.elems, elem)

	return nil
}

func (q *Queue[T]) EnqueueMany(elems []T) error {
	if len(elems) == 0 {
		return nil
	} else if q == nil {
		return gers.ErrNilReceiver
	}

	q.elems = append(q.elems, elems...)

	return nil
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q == nil || len(q.elems) == 0 {
		return *new(T), ErrEmptyQueue
	}

	elem := q.elems[0]
	q.elems = q.elems[1:]

	return elem, nil
}
