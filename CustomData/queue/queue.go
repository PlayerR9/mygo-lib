package queue

import "github.com/PlayerR9/mygo-lib/common"

// Queue is an interface defining a queue.
type Queue[T any] interface {
	// Enqueue adds an element at the end of the queue.
	//
	// Parameters:
	//   - elem: The element to add.
	//
	// Returns:
	//   - error: An error if the element could not be added to the queue.
	//
	// Errors:
	// 	- common.ErrNilReceiver: If the receiver is nil.
	// 	- ErrFullQueue: If the queue has a capacity and that capacity has been reached.
	//    - any other error that may be returned.
	Enqueue(elem T) error

	// Dequeue removes the first element from the queue.
	//
	// Returns:
	//   - T: The element that was removed.
	//   - error: An error if the element could not be removed from the queue.
	//
	// Errors:
	//   - ErrEmptyQueue: If the queue is empty.
	//   - common.ErrNilReceiver: If the receiver is nil.
	//   - any other error that may be returned.
	Dequeue() (T, error)

	// Front returns the element at the start of the queue without removing it.
	//
	// Returns:
	//   - T: The element at the start of the queue.
	//   - error: An error if the element could not be peeked from the queue.
	//
	// Errors:
	//   - common.ErrNilReceiver: If the receiver is nil.
	//   - common.ErrEmptyQueue: If the queue is empty.
	//   - any other error that may be returned.
	Front() (T, error)

	// Size returns the number of elements in the queue.
	//
	// Returns:
	//   - uint: The number of elements in the queue.
	//
	// If the receiver is nil, then 0 is returned.
	Size() uint

	// IsEmpty checks whether the queue is empty.
	//
	// Returns:
	//   - bool: True if the queue is empty, false otherwise.
	//
	// If the receiver is nil, then true is returned.
	IsEmpty() bool
}

// Enqueue adds multiple elements to the queue in the order they are passed. If the queue implements
// the `EnqueueMany` method, then that method is used instead.
//
// Parameters:
//   - queue: The queue to which the elements are added.
//   - elems: Variadic parameters representing the elements to be added.
//
// Returns:
//   - uint: The number of elements successfully enqueued onto the queue.
//   - error: An error if the queue is nil or if there is an issue enqueuing one of the elements.
//
// Errors:
//   - common.ErrNilParam: If the queue is nil.
//   - ErrFullQueue: If not all elements could be enqueued onto the queue.
//   - any error returned by the `Enqueue()` method of the queue.
func Enqueue[T any](queue Queue[T], elems ...T) (uint, error) {
	lenElems := uint(len(elems))
	if lenElems == 0 {
		return 0, nil
	} else if queue == nil {
		return 0, common.NewErrNilParam("queue")
	}

	q, ok := queue.(interface{ EnqueueMany(elems []T) (uint, error) })
	if ok {
		n, err := q.EnqueueMany(elems)
		return n, err
	}

	for i, elem := range elems {
		err := queue.Enqueue(elem)
		if err != nil {
			return uint(i), err
		}
	}

	return lenElems, nil
}

// Free frees the queue. If the queue implements `Type` interface, then its `Free()`
// method is called. If not, then the queue is cleared by dequeueing all elements from the queue.
//
// Parameters:
//   - queue: The queue to free.
func Free[T any](queue Queue[T]) {
	if queue == nil {
		return
	}

	if q, ok := queue.(common.Freeable); ok {
		q.Free()
		return
	}

	for {
		_, err := queue.Dequeue()
		if err != nil {
			break
		}
	}
}

// Reset resets the queue for reuse. If the queue implements `Resetter` interface,
// then its `Reset()` method is called. If not, then the queue is cleared by dequeueing all
// elements from the queue.
//
// Parameters:
//   - queue: The queue to reset.
func Reset[T any](queue Queue[T]) {
	if queue == nil || common.Reset(queue) {
		return
	}

	for {
		_, err := queue.Dequeue()
		if err != nil {
			break
		}
	}
}
