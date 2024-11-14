package internal

import (
	"sync"

	"github.com/PlayerR9/mygo-lib/common"
	"github.com/PlayerR9/mygo-lib/mem"
)

// StackNode is a node in the stack.
type StackNode[T any] struct {
	// elem is the element.
	elem T

	// prev is the previous node.
	prev *mem.Ref[*StackNode[T]]

	// mu is the mutex.
	mu sync.RWMutex
}

// free is a private method that frees the stack node.
func (node *StackNode[T]) free() error {
	if node == nil {
		return mem.ErrNilReceiver
	}

	node.mu.Lock()
	defer node.mu.Unlock()

	if node.prev == nil {
		node.elem = *new(T)

		return nil
	}

	err := mem.Free("node.prev", node.prev)
	if err != nil {
		return err
	}

	node.prev = nil
	node.elem = *new(T)

	return nil
}

// NewStackNode creates a new stack node with the given element.
//
// Parameters:
//   - elem: The element of the node.
//
// Returns:
//   - *mem.Ref[*StackNode[T]]: The new node. Never returns nil.
func NewStackNode[T any](elem T) *mem.Ref[*StackNode[T]] {
	node := &StackNode[T]{
		elem: elem,
		prev: nil,
	}

	ref := mem.New(node, node.free)
	return ref
}

// SetPrev sets the previous node of the stack node. This also transfers
// ownership of `prev` node to the receiver.
//
// Parameters:
//   - prev: The node to set as the previous node.
//
// Returns:
//   - error: An error of type common.ErrNilReceiver if the receiver is nil.
func (n *StackNode[T]) SetPrev(prev *mem.Ref[*StackNode[T]]) error {
	if n == nil {
		return common.ErrNilReceiver
	}

	n.mu.Lock()
	defer n.mu.Unlock()

	if n.prev != nil {
		// Free the previous node.
		err := mem.Free("n.prev", n.prev)
		if err != nil {
			return err
		}
	}

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
