package internal

import (
	"sync"

	"github.com/PlayerR9/mygo-lib/common"
)

// link_node links the given nodes together, where `a` is the previous node and
// `b` is the next node. If `a` is `nil`, this will set `b.prev` to `nil`.
// If `b` is `nil`, this will set `a.next` to `nil`. If both are `nil`, this
// will do nothing.
//
// Parameters:
//   - a: The previous node.
//   - b: The next node.
func link_node[T any](a, b *ListNode[T]) {
	if a == nil && b == nil {
		return
	}

	if a == nil {
		b.mu.Lock()
		defer b.mu.Unlock()

		b.prev = nil

		return
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	if b != nil {
		b.mu.Lock()
		defer b.mu.Unlock()

		b.prev = a
	}

	a.next = b
}

// ListNode is a node in the list.
type ListNode[T any] struct {
	// elem is the element.
	elem T

	// next is the next node.
	next *ListNode[T]

	// prev is the previous node.
	prev *ListNode[T]

	// mu is the mutex.
	mu sync.RWMutex
}

// Free implements common.Type.
func (node *ListNode[T]) Free() {
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

	if node.prev != nil {
		node.prev.Free()
		node.prev = nil
	}
}

// NewListNode creates a new list node with the given element.
//
// Parameters:
//   - elem: The element of the node.
//
// Returns:
//   - *ListNode[T]: The new node. Never returns nil.
func NewListNode[T any](elem T) *ListNode[T] {
	return &ListNode[T]{
		elem: elem,
		next: nil,
		prev: nil,
	}
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

	link_node(n, next)

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

	link_node(prev, n)

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
