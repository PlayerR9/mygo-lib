package internal

import (
	"sync"

	"github.com/PlayerR9/mygo-lib/common"
)

// StackNode is a node in the stack.
type StackNode[T any] struct {
	// elem is the element.
	elem T

	// prev is the previous node.
	prev *StackNode[T]

	// mu is the mutex.
	mu sync.RWMutex
}

// Free implements common.Type.
func (node *StackNode[T]) Free() {
	if node == nil {
		return
	}

	node.mu.Lock()
	defer node.mu.Unlock()

	common.Free(node.prev)

	node.elem = *new(T)
	node.prev = nil
}

// NewStackNode creates a new stack node with the given element.
//
// Parameters:
//   - elem: The element of the node.
//
// Returns:
//   - *StackNode[T]: The new node. Never returns nil.
func NewStackNode[T any](elem T) *StackNode[T] {
	return &StackNode[T]{
		elem: elem,
	}
}

// SetPrev sets the previous node of the stack node.
//
// Parameters:
//   - prev: The node to set as the previous node.
//
// Returns:
//   - error: An error if the receiver is nil.
//
// Errors:
//   - common.ErrNilReceiver: If the receiver is nil.
func (n *StackNode[T]) SetPrev(prev *StackNode[T]) error {
	if n == nil {
		return common.ErrNilReceiver
	}

	n.mu.Lock()
	defer n.mu.Unlock()

	n.prev = prev

	return nil
}

// MustGetPrev returns the previous node of the stack node.
//
// Returns:
//   - *StackNode[T]: The previous node.
//
// Panics:
//   - common.ErrNilReceiver: If the receiver is nil.
func (n *StackNode[T]) MustGetPrev() *StackNode[T] {
	if n == nil {
		panic(common.ErrNilReceiver)
	}

	n.mu.RLock()
	defer n.mu.RUnlock()

	return n.prev
}

// MustGetElem returns the element of the stack node.
//
// Returns:
//   - T: The element.
//
// Panics:
//   - common.ErrNilReceiver: If the receiver is nil.
func (n *StackNode[T]) MustGetElem() T {
	if n == nil {
		panic(common.ErrNilReceiver)
	}

	n.mu.RLock()
	defer n.mu.RUnlock()

	return n.elem
}
