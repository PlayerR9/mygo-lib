package queue

import (
	"sync"

	"github.com/PlayerR9/mygo-lib/CustomData/queue/internal"
	"github.com/PlayerR9/mygo-lib/common"
)

// LinkedQueue is a simple implementation of a queue that is backed by an linked list.
// This implementation is thread-safe.
//
// An empty linked queue can be created using the `queue := new(queue.LinkedQueue[T])` constructor.
type LinkedQueue[T any] struct {
	// front is the front of the queue.
	front *internal.QueueNode[T]

	// back is the back of the queue.
	back *internal.QueueNode[T]

	// front_mu is the mutex for the front of the queue.
	front_mu sync.RWMutex

	// back_mu is the mutex for the back of the queue.
	back_mu sync.Mutex
}

// Enqueue implements Queue.
//
// Never returns ErrFullQueue.
func (q *LinkedQueue[T]) Enqueue(elem T) error {
	if q == nil {
		return common.ErrNilReceiver
	}

	q.back_mu.Lock()
	defer q.back_mu.Unlock()

	node := internal.NewQueueNode(elem)

	if q.back == nil {
		q.front_mu.Lock()
		defer q.front_mu.Unlock()

		q.back = node
		q.front = node

		return nil
	}

	_ = q.back.SetNext(node)
	q.back = node

	return nil
}

// Dequeue implements Queue.
func (q *LinkedQueue[T]) Dequeue() (T, error) {
	if q == nil {
		return *new(T), common.ErrNilReceiver
	}

	q.front_mu.Lock()
	defer q.front_mu.Unlock()

	if q.front == nil {
		return *new(T), ErrEmptyQueue
	}

	front := q.front
	q.front = front.MustGetNext()
	_ = front.SetNext(nil) // Clear the reference.

	if q.front == nil {
		q.back_mu.Lock()
		q.back = nil
		q.back_mu.Unlock()
	}

	elem := front.MustGetElem()
	return elem, nil
}

// Front implements Queue.
func (q *LinkedQueue[T]) Front() (T, error) {
	if q == nil {
		return *new(T), common.ErrNilReceiver
	}

	q.front_mu.RLock()
	defer q.front_mu.RUnlock()

	if q.front == nil {
		return *new(T), ErrEmptyQueue
	}

	elem := q.front.MustGetElem()
	return elem, nil
}

// IsEmpty implements Queue.
func (q *LinkedQueue[T]) IsEmpty() bool {
	if q == nil {
		return true
	}

	q.front_mu.RLock()
	defer q.front_mu.RUnlock()

	return q.front == nil
}

// Size implements Queue.
func (q *LinkedQueue[T]) Size() uint {
	if q == nil {
		return 0
	}

	q.front_mu.RLock()
	defer q.front_mu.RUnlock()

	q.back_mu.Lock()
	defer q.back_mu.Unlock()

	var size uint

	for c := q.front; c != nil; c = c.MustGetNext() {
		size++
	}

	return size
}

// Free implements common.Typer.
func (q *LinkedQueue[T]) Free() {
	if q == nil {
		return
	}

	q.front_mu.Lock()
	defer q.front_mu.Unlock()

	q.back_mu.Lock()
	defer q.back_mu.Unlock()

	if q.front != nil {
		q.front.Free()
		q.front = nil
	}

	q.back = nil
}

// Reset implements common.Resetter.
func (q *LinkedQueue[T]) Reset() {
	if q == nil {
		return
	}

	q.front_mu.Lock()
	defer q.front_mu.Unlock()

	q.back_mu.Lock()
	defer q.back_mu.Unlock()

	if q.front != nil {
		q.front.Free()
		q.front = nil
	}

	q.back = nil
}

// link_elements creates a slice of *QueueNode from the given elements,
// linking each node to the next node in the slice. The last node in the
// slice will have no previous node.
//
// Parameters:
//   - elems: A slice of elements to be converted into queue nodes.
//
// Returns:
//   - []*QueueNode[T]: A slice of linked queue nodes, where each node points
//     to the next node as its previous node.
func link_elements[T any](elems []T) []*internal.QueueNode[T] {
	slice := make([]*internal.QueueNode[T], 0, len(elems))

	for _, elem := range elems {
		node := internal.NewQueueNode(elem)
		slice = append(slice, node)
	}

	for i := range slice[:len(slice)-1] {
		_ = slice[i].SetNext(slice[i+1])
	}

	return slice
}

// EnqueueMany adds multiple elements to the queue in the order they are passed.
//
// Parameters:
//   - elems: A slice of elements to be added to the queue.
//
// Returns:
//   - uint: The number of elements successfully enqueued onto the queue.
//   - error: An error of type common.ErrNilReceiver if the receiver is nil.
func (q *LinkedQueue[T]) EnqueueMany(elems []T) (uint, error) {
	lenElems := uint(len(elems))
	if lenElems == 0 {
		return 0, nil
	} else if q == nil {
		return 0, common.ErrNilReceiver
	}

	q.back_mu.Lock()
	defer q.back_mu.Unlock()

	slice := link_elements(elems)

	if q.back == nil {
		q.front_mu.Lock()
		defer q.front_mu.Unlock()

		q.front = slice[0]
	} else {
		_ = q.back.SetNext(slice[0])
	}

	q.back = slice[len(slice)-1]

	clear(slice)

	return lenElems, nil
}
