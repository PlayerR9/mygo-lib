package stack

import (
	"slices"
	"sync"

	"github.com/PlayerR9/mygo-lib/common"
)

// ArrayStack is a simple implementation of a stack that is backed by an array.
// This implementation is thread-safe.
//
// An empty array stack can be created using the `stack := new(ArrayStack[T])` constructor.
type ArrayStack[T any] struct {
	// slice is the backing array.
	slice []T

	// lenSlice is the number of elements in the slice.
	lenSlice uint

	// mu is the mutex.
	mu sync.RWMutex
}

// Push implements Stack.
//
// Never returns ErrFullStack.
func (s *ArrayStack[T]) Push(elem T) error {
	if s == nil {
		return common.ErrNilReceiver
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.slice = append(s.slice, elem)
	s.lenSlice++

	return nil
}

// Pop implements Stack.
func (s *ArrayStack[T]) Pop() (T, error) {
	if s == nil {
		return *new(T), common.ErrNilReceiver
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.lenSlice == 0 {
		return *new(T), ErrEmptyStack
	}

	top := s.slice[s.lenSlice-1]
	s.slice = s.slice[:s.lenSlice-1]
	s.lenSlice--

	return top, nil
}

// Peek implements Stack.
func (s *ArrayStack[T]) Peek() (T, error) {
	if s == nil {
		return *new(T), common.ErrNilReceiver
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.lenSlice == 0 {
		return *new(T), ErrEmptyStack
	}

	return s.slice[s.lenSlice-1], nil
}

// IsEmpty implements Stack.
func (s *ArrayStack[T]) IsEmpty() bool {
	if s == nil {
		return true
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.lenSlice == 0
}

// Size implements Stack.
func (s *ArrayStack[T]) Size() uint {
	if s == nil {
		return 0
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.lenSlice
}

// Free implements common.Typer.
func (s *ArrayStack[T]) Free() {
	if s == nil {
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	clear(s.slice)
	s.slice = nil
	s.lenSlice = 0
}

// Reset implements common.Resetter.
func (s *ArrayStack[T]) Reset() {
	if s == nil {
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	clear(s.slice)
	s.slice = nil
	s.lenSlice = 0
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
func (s *ArrayStack[T]) PushMany(elems []T) (uint, error) {
	lenElems := uint(len(elems))
	if lenElems == 0 {
		return 0, nil
	} else if s == nil {
		return 0, common.ErrNilReceiver
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	slice := make([]T, len(elems))
	copy(slice, elems)
	slices.Reverse(slice)

	s.slice = append(s.slice, slice...)
	s.lenSlice += lenElems

	clear(slice)
	slice = nil

	return lenElems, nil
}
