package stack

import (
	"sync"

	"github.com/PlayerR9/mygo-lib/CustomData/stack/internal"
	"github.com/PlayerR9/mygo-lib/common"
)

// LinkedStack is a simple implementation of a stack that is backed by an linked list.
// This implementation is thread-safe.
//
// An empty linked stack can be created using the `stack := new(stack.LinkedStack[T])` constructor
// or the provided `NewLinkedStack` function.
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

	node := internal.NewStackNode(elem)

	s.mu.Lock()
	defer s.mu.Unlock()

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

// Free implements common.Type.
func (s *LinkedStack[T]) Free() {
	if s == nil {
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	common.Free(s.front)

	s.front = nil
}

// Reset implements common.Resetter.
func (s *LinkedStack[T]) Reset() {
	if s == nil {
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	common.Free(s.front)

	s.front = nil
}

// NewLinkedStack creates a new stack from a slice.
//
// Parameters:
//   - elems: The elements to add to the stack.
//
// Returns:
//   - *LinkedStack[T]: The new stack. Never returns nil.
func NewLinkedStack[T any](elems ...T) *LinkedStack[T] {
	stack := new(LinkedStack[T])
	if len(elems) == 0 {
		return stack
	}

	_, _ = stack.PushMany(elems)
	return stack
}

// PushMany adds multiple elements to the stack in reverse order.
//
// Parameters:
//   - elems: A slice of elements to be added to the stack.
//
// Returns:
//   - uint: The number of elements successfully pushed onto the stack.
//   - error: An error of type common.ErrNilReceiver if the receiver is nil.
//
// The elements are pushed onto the stack in reverse order, meaning the last element
// in the slice will be at the top of the stack after the operation. If the slice
// is empty, the function returns immediately with zero elements pushed.
func (s *LinkedStack[T]) PushMany(elems []T) (uint, error) {
	lenElems := uint(len(elems))
	if lenElems == 0 {
		return 0, nil
	} else if s == nil {
		return 0, common.ErrNilReceiver
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	for i := len(elems) - 1; i >= 0; i-- {
		node := internal.NewStackNode(elems[i])
		_ = node.SetPrev(s.front)

		s.front = node
	}

	return lenElems, nil
}
