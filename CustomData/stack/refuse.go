package stack

import (
	"slices"
	"sync"

	"github.com/PlayerR9/mygo-lib/common"
)

// Refuse is a wrapper for a Stack that allows elements to be refused, meaning that
// once popped we can decide whether to accept that popping sequence or not.
type Refuse[T any] struct {
	// stack is the internal stack.
	stack Stack[T]

	// popped is the stack of elements that were popped.
	popped []T

	// mu is the mutex for the RefusableStack.
	mu sync.RWMutex
}

// Size implements the Stack interface.
//
// Panics:
//   - common.ErrInvalidObject: If the method `Free()` is called.
func (r *Refuse[T]) Size() uint {
	if r == nil {
		return 0
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.stack == nil {
		panic(common.NewErrInvalidObject("Size"))
	}

	size := r.stack.Size()
	return size
}

// IsEmpty implements the Stack interface.
//
// Panics:
//   - common.ErrInvalidObject: If the method `Free()` is called.
func (r *Refuse[T]) IsEmpty() bool {
	if r == nil {
		return true
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.stack == nil {
		panic(common.NewErrInvalidObject("IsEmpty"))
	}

	ok := r.stack.IsEmpty()
	return ok
}

// Push implements the Stack interface.
//
// Errors:
//   - common.ErrInvalidObject: If the method `Free()` is called.
func (r *Refuse[T]) Push(elem T) error {
	if r == nil {
		return common.ErrNilReceiver
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if r.stack == nil {
		return common.NewErrInvalidObject("Push")
	}

	err := r.stack.Push(elem)
	if err != nil {
		return err
	}

	return nil
}

// Pop implements the Stack interface.
//
// Errors:
//   - common.ErrInvalidObject: If the method `Free()` is called.
func (r *Refuse[T]) Pop() (T, error) {
	if r == nil {
		return *new(T), common.ErrNilReceiver
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if r.stack == nil {
		return *new(T), common.NewErrInvalidObject("Pop")
	}

	top, err := r.stack.Pop()
	if err != nil {
		return *new(T), err
	}

	r.popped = append(r.popped, top)

	return top, nil
}

// Peek implements the Stack interface.
//
// Errors:
//   - common.ErrInvalidObject: If the method `Free()` is called.
func (r *Refuse[T]) Peek() (T, error) {
	if r == nil {
		return *new(T), common.ErrNilReceiver
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.stack == nil {
		return *new(T), common.NewErrInvalidObject("Peek")
	}

	top, err := r.stack.Peek()
	if err != nil {
		return *new(T), err
	}

	return top, nil
}

// Free implements common.Type.
func (r *Refuse[T]) Free() {
	if r == nil {
		return
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	Free(r.stack)
	r.stack = nil

	clear(r.popped)
	r.popped = nil
}

// Reset implements common.Resetter.
func (r *Refuse[T]) Reset() {
	if r == nil {
		return
	}

	r.mu.Lock()
	defer r.mu.Lock()

	Reset(r.stack)

	clear(r.popped)
	r.popped = nil
}

// NewRefuse creates a new stack from a slice.
//
// Parameters:
//   - stack: The stack that can be refused.
//
// Returns:
//   - *Refuse[T]: The new stack. Nil if the stack is nil.
func NewRefuse[T any](stack Stack[T]) *Refuse[T] {
	if stack == nil {
		return nil
	}

	return &Refuse[T]{
		stack: stack,
	}
}

// Accept accepts all the elements that were popped. It's a no-op if no element was popped.
//
// Returns:
//   - error: An error of type common.ErrNilReceiver if the receiver is nil.
func (r *Refuse[T]) Accept() error {
	if r == nil {
		return common.ErrNilReceiver
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	clear(r.popped)
	r.popped = nil

	return nil
}

// Refuse refuses any element that was popped since the last time `Accept()` was called.
// It's a no-op if no element was popped.
//
// Returns:
//   - error: An error if refusing the element failed.
//
// Errors:
//   - common.ErrNilReceiver: If the receiver is nil.
//   - common.ErrInvalidObject: If the method `Free()` is called.
func (r *Refuse[T]) Refuse() error {
	if r == nil {
		return common.ErrNilReceiver
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if r.stack == nil {
		return common.NewErrInvalidObject("Refuse")
	}

	var i int
	var err error

	for i = len(r.popped) - 1; i >= 0 && err == nil; i-- {
		err = r.stack.Push(r.popped[i])
	}

	if i == 0 {
		clear(r.popped)
		r.popped = nil
	} else {
		clear(r.popped[i:])
		r.popped = r.popped[:i]
	}

	return err
}

// RefuseOne refuses the last popped element. It's a no-op if no element was popped.
//
// Returns:
//   - error: An error if refusing the element failed.
//
// Errors:
//   - common.ErrNilReceiver: If the receiver is nil.
//   - common.ErrInvalidObject: If the method `Free()` is called.
func (r *Refuse[T]) RefuseOne() error {
	if r == nil {
		return common.ErrNilReceiver
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if len(r.popped) == 0 {
		return nil
	} else if r.stack == nil {
		return common.NewErrInvalidObject("RefuseOne")
	}

	top := r.popped[len(r.popped)-1]
	r.popped = r.popped[:len(r.popped)-1]

	err := r.stack.Push(top)
	return err
}

// Popped returns the elements that were popped from the stack since the last
// `Accept()` or `Refuse()` operation. The returned slice contains the elements
// in the order they were popped, with the most recently popped element at the
// first position.
//
// Returns:
//   - []T: The elements that were popped.
func (r *Refuse[T]) Popped() []T {
	if r == nil {
		return nil
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	if len(r.popped) == 0 {
		return nil
	}

	slice := make([]T, len(r.popped))
	copy(slice, r.popped)

	slices.Reverse(slice)

	return slice
}

// IsValid checks whether the Refuse stack is in a valid state.
//
// Returns:
//   - bool: True if the stack is valid and not nil, false otherwise.
func (r *Refuse[T]) IsValid() bool {
	if r == nil {
		return false
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.stack != nil
}
