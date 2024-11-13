package stack

import (
	"sync"

	"github.com/PlayerR9/mygo-lib/CustomData/stack/internal"
	"github.com/PlayerR9/mygo-lib/common"
)

// LinkedStack is a simple implementation of a stack that is backed by an linked list.
// This implementation is thread-safe.
//
// An empty linked stack can be created using the `stack := new(stack.LinkedStack[T])` constructor.
type LinkedStack[T any] struct {
	// front is the front of the stack.
	front *internal.StackNode[T]

	// mu is the mutex.
	mu sync.RWMutex
}

// Push implements Stack.
//
// Never returns ErrFullStack.
func (s *LinkedStack[T]) Push(elem T) error {
	if s == nil {
		return common.ErrNilReceiver
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	node := internal.NewStackNode(elem)

	if s.front == nil {
		s.front = node

		return nil
	}

	_ = node.SetPrev(s.front)
	s.front = node

	return nil
}

// Pop implements Stack.
func (s *LinkedStack[T]) Pop() (T, error) {
	if s == nil {
		return *new(T), common.ErrNilReceiver
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.front == nil {
		return *new(T), ErrEmptyStack
	}

	top := s.front
	s.front = top.MustGetPrev()

	_ = top.SetPrev(nil) // Clear the reference.

	elem := top.MustGetElem()
	return elem, nil
}

// Peek implements Stack.
func (s *LinkedStack[T]) Peek() (T, error) {
	if s == nil {
		return *new(T), common.ErrNilReceiver
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.front == nil {
		return *new(T), ErrEmptyStack
	}

	elem := s.front.MustGetElem()
	return elem, nil
}

// IsEmpty implements Stack.
func (s *LinkedStack[T]) IsEmpty() bool {
	if s == nil {
		return true
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.front == nil
}

// Size implements Stack.
func (s *LinkedStack[T]) Size() uint {
	if s == nil {
		return 0
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	var size uint

	for c := s.front; c != nil; c = c.MustGetPrev() {
		size++
	}

	return size
}

// Free implements common.Typer.
func (s *LinkedStack[T]) Free() {
	if s == nil {
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.front != nil {
		s.front.Free()
		s.front = nil
	}
}

// Reset implements common.Resetter.
func (s *LinkedStack[T]) Reset() {
	if s == nil {
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.front != nil {
		s.front.Free()
		s.front = nil
	}
}

// link_elements creates a slice of StackNode pointers from the given elements,
// linking each node to the next node in the slice. The last node in the slice
// will have no previous node.
//
// Parameters:
//   - elems: A slice of elements to be converted into stack nodes.
//
// Returns:
//   - []*internal.StackNode[T]: A slice of linked stack nodes, where each node
//     points to the next node as its previous node.
func link_elements[T any](elems []T) []*internal.StackNode[T] {
	slice := make([]*internal.StackNode[T], 0, len(elems))

	for _, elem := range elems {
		node := internal.NewStackNode(elem)
		slice = append(slice, node)
	}

	for i := range slice[:len(slice)-1] {
		_ = slice[i].SetPrev(slice[i+1])
	}

	return slice
}

// PushMany adds multiple elements to the stack in reverse order, meaning that the
// first element in the slice will be at the top of the stack after the operation.
//
// Parameters:
//   - elems: A slice of elements to be added to the stack.
//
// Returns:
//   - uint: The number of elements successfully pushed onto the stack.
//   - error: An error of type common.ErrNilReceiver if the receiver is nil.
func (s *LinkedStack[T]) PushMany(elems []T) (uint, error) {
	lenElems := uint(len(elems))
	if lenElems == 0 {
		return 0, nil
	} else if s == nil {
		return 0, common.ErrNilReceiver
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	slice := link_elements(elems)

	slice[len(slice)-1].SetPrev(s.front)
	s.front = slice[0]

	clear(slice)

	return lenElems, nil
}
