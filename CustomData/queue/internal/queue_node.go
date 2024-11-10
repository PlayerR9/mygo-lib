package internal

import (
	"sync"

	"github.com/PlayerR9/mygo-lib/common"
)

// QueueNode is a node in the queue.
type QueueNode[T any] struct {
	// elem is the element.
	elem T

	// next is the next node.
	next *QueueNode[T]

	// mu is the mutex.
	mu sync.RWMutex
}

// Free implements common.Type.
func (node *QueueNode[T]) Free() {
	if node == nil {
		return
	}

	node.mu.Lock()
	defer node.mu.Unlock()

	node.elem = *new(T)

	if node.next != nil {
		node.next.Free()
		node.next = nil
	}
}

// NewQueueNode creates a new queue node with the given element.
//
// Parameters:
//   - elem: The element of the node.
//
// Returns:
//   - *QueueNode[T]: The new node. Never returns nil.
func NewQueueNode[T any](elem T) *QueueNode[T] {
	return &QueueNode[T]{
		elem: elem,
		next: nil,
	}
}

// SetNext sets the next node of the queue node.
//
// Parameters:
//   - next: The node to set as the next node.
//
// Returns:
//   - error: An error of type common.ErrNilReceiver if the receiver is nil.
func (n *QueueNode[T]) SetNext(next *QueueNode[T]) error {
	if n == nil {
		return common.ErrNilReceiver
	}

	n.mu.Lock()
	defer n.mu.Unlock()

	n.next = next

	return nil
}

// MustGetNext returns the next node of the queue node.
//
// Returns:
//   - *QueueNode[T]: The next node.
//
// Panics:
//   - common.ErrNilReceiver: If the receiver is nil.
func (n *QueueNode[T]) MustGetNext() *QueueNode[T] {
	if n == nil {
		panic(common.ErrNilReceiver)
	}

	n.mu.RLock()
	defer n.mu.RUnlock()

	return n.next
}

// MustGetElem returns the element of the queue node.
//
// Returns:
//   - T: The element.
//
// Panics:
//   - common.ErrNilReceiver: If the receiver is nil.
func (n *QueueNode[T]) MustGetElem() T {
	if n == nil {
		panic(common.ErrNilReceiver)
	}

	n.mu.RLock()
	defer n.mu.RUnlock()

	return n.elem
}
