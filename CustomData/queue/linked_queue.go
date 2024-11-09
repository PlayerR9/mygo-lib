package queue

import (
	"sync"

	"github.com/PlayerR9/mygo-lib/common"
)

// LinkedQueue is a simple implementation of a queue that is backed by an linked list.
// This implementation is thread-safe.
type LinkedQueue[T any] struct {
	// front is the front of the queue.
	front *QueueNode[T]

	// front_mu is the mutex.
	front_mu sync.RWMutex

	// back is the back of the queue.
	back *QueueNode[T]

	// back_mu is the mutex.
	back_mu sync.RWMutex
}

// Enqueue implements Queue.
//
// Never returns ErrFullQueue.
func (s *LinkedQueue[T]) Enqueue(elem T) error {
	if s == nil {
		return common.ErrNilReceiver
	}

	node := &QueueNode[T]{
		elem: elem,
	}

	s.back_mu.Lock()
	defer s.back_mu.Unlock()

	if s.back == nil {
		s.front_mu.Lock()
		s.front = node
		s.front_mu.Unlock()
	} else {
		s.back.next = node
	}

	s.back = node

	return nil
}

// Dequeue implements Queue.
func (s *LinkedQueue[T]) Dequeue() (T, error) {
	if s == nil {
		return *new(T), common.ErrNilReceiver
	}

	s.front_mu.RLock()
	is_empty := s.front == nil
	s.front_mu.RUnlock()

	if is_empty {
		return *new(T), ErrEmptyQueue
	}

	s.front_mu.Lock()
	defer s.front_mu.Unlock()

	top := s.front
	s.front = top.next

	if s.front == nil {
		s.back_mu.Lock()
		s.back = nil
		s.back_mu.Unlock()
	}

	top.next = nil // Clear the reference.

	return top.elem, nil
}

// Front implements Queue.
func (s *LinkedQueue[T]) Front() (T, error) {
	if s == nil {
		return *new(T), common.ErrNilReceiver
	}

	s.front_mu.RLock()
	front := s.front
	s.front_mu.RUnlock()

	if front == nil {
		return *new(T), ErrEmptyQueue
	} else {
		return front.elem, nil
	}
}

// IsEmpty implements Queue.
func (s *LinkedQueue[T]) IsEmpty() bool {
	if s == nil {
		return true
	}

	s.front_mu.RLock()
	defer s.front_mu.RUnlock()

	return s.front == nil
}

// Size implements Queue.
func (s *LinkedQueue[T]) Size() uint {
	if s == nil {
		return 0
	}

	s.front_mu.RLock()
	defer s.front_mu.RUnlock()

	var size uint

	for c := s.front; c != nil; c = c.next {
		size++
	}

	return size
}

// Free implements common.Type.
func (s *LinkedQueue[T]) Free() {
	if s == nil {
		return
	}

	s.front_mu.Lock()
	defer s.front_mu.Unlock()

	if s.front != nil {
		s.front.Free()
	}

	s.front = nil

	s.back_mu.Lock()
	defer s.back_mu.Unlock()

	s.back = nil
}

// NewLinkedQueue creates a new queue from a slice.
//
// Parameters:
//   - elems: The elements to add to the queue.
//
// Returns:
//   - *LinkedQueue[T]: The new queue. Never returns nil.
func NewLinkedQueue[T any](elems ...T) *LinkedQueue[T] {
	if len(elems) == 0 {
		return &LinkedQueue[T]{
			front: nil,
			back:  nil,
		}
	}

	first_node := &QueueNode[T]{
		elem: elems[0],
	}

	last_node := first_node

	for _, elem := range elems[1:] {
		node := &QueueNode[T]{
			elem: elem,
		}

		last_node.next = node
		last_node = node
	}

	return &LinkedQueue[T]{
		front: first_node,
		back:  last_node,
	}
}

// Reset resets the queue for reuse. Does nothing if the receiver is nil.
func (s *LinkedQueue[T]) Reset() {
	if s == nil {
		return
	}

	s.front_mu.Lock()
	defer s.front_mu.Unlock()

	if s.front != nil {
		s.front.Free()
	}

	s.front = nil

	s.back_mu.Lock()
	defer s.back_mu.Unlock()

	s.back = nil
}
