package internal

import (
	"sync"

	"github.com/PlayerR9/mygo-lib/common"
	"github.com/PlayerR9/mygo-lib/mem"
)

// ListNode is a node in the list.
type ListNode[T any] struct {
	// elem is the element.
	elem T

	// next is the next node.
	next *ListNode[T]

	// prev is the previous node.
	prev *mem.Ref[*ListNode[T]]

	// mu is the mutex.
	mu sync.RWMutex
}

// free is a private method that frees the list node.
func (node *ListNode[T]) free() error {
	if node == nil {
		return mem.ErrNilReceiver
	}

	node.mu.Lock()
	defer node.mu.Unlock()

	node.elem = *new(T)

	if node.prev == nil {
		node.next = nil

		return nil
	}

	err := mem.Free("node.prev", node.prev)
	if err != nil {
		return err
	}

	node.prev = nil
	node.next = nil

	return nil
}

// NewListNode creates a new list node with the given element.
//
// Parameters:
//   - elem: The element of the node.
//
// Returns:
//   - *mem.Ref[*ListNode[T]]: A reference to the list node. Never returns nil.
func NewListNode[T any](elem T) *mem.Ref[*ListNode[T]] {
	node := &ListNode[T]{
		elem: elem,
		next: nil,
		prev: nil,
	}

	ref := mem.New(node, node.free)
	return ref
}

// SetNext sets the next node of the list node.
//
// Parameters:
//   - next: The node to set as the next node.
//
// Returns:
//   - error: An error of type common.ErrNilReceiver if the receiver is nil.
func (n *ListNode[T]) SetNext(next *ListNode[T]) error {
	if n == nil {
		return common.ErrNilReceiver
	}

	n.mu.Lock()
	defer n.mu.Unlock()

	if next == nil {
		n.next = nil

		return nil
	}

	next.mu.Lock()
	defer next.mu.Unlock()

	next.prev = n

	n.next = next

	return nil
}

// SetPrev sets the previous node of the list node.
//
// Parameters:
//   - prev: The node to set as the previous node.
//
// Returns:
//   - error: An error of type common.ErrNilReceiver if the receiver is nil.
func (n *ListNode[T]) SetPrev(prev *ListNode[T]) error {
	if n == nil {
		return common.ErrNilReceiver
	}

	n.mu.Lock()
	defer n.mu.Unlock()

	if prev == nil {
		err := mem.Free("n.prev", n.prev)
		if err != nil {
			return err
		}

		n.prev = nil

		return nil
	}

	n.mu.Lock()
	defer n.mu.Unlock()

	n.prev = prev

	prev.next = n

	return nil
}

// MustGetNext returns the next node of the list node.
//
// Returns:
//   - *ListNode[T]: The next node.
//
// Panics:
//   - common.ErrNilReceiver: If the receiver is nil.
func (n *ListNode[T]) MustGetNext() *ListNode[T] {
	if n == nil {
		panic(common.ErrNilReceiver)
	}

	n.mu.RLock()
	defer n.mu.RUnlock()

	return n.next
}

// MustGetPrev returns the previous node of the list node.
//
// Returns:
//   - *ListNode[T]: The previous node.
//
// Panics:
//   - common.ErrNilReceiver: If the receiver is nil.
func (n *ListNode[T]) MustGetPrev() *ListNode[T] {
	if n == nil {
		panic(common.ErrNilReceiver)
	}

	n.mu.RLock()
	defer n.mu.RUnlock()

	return n.prev
}

// MustGetElem returns the element of the list node.
//
// Returns:
//   - T: The element.
//
// Panics:
//   - common.ErrNilReceiver: If the receiver is nil.
func (n *ListNode[T]) MustGetElem() T {
	if n == nil {
		panic(common.ErrNilReceiver)
	}

	n.mu.RLock()
	defer n.mu.RUnlock()

	return n.elem
}
