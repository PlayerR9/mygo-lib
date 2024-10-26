package rws

import (
	"sync"

	"github.com/PlayerR9/mygo-lib/common"
)

// Var is a variable that can be read and written safely.
type Var[T any] struct {
	// data is the value of the variable.
	data T

	// mu is the mutex for the variable.
	mu sync.RWMutex
}

// New creates a new Var.
//
// Parameters:
//   - data: The initial value of the variable.
//
// Returns:
//   - *Var[T]: The new Var. Never returns nil.
func New[T any](data T) *Var[T] {
	return &Var[T]{
		data: data,
	}
}

// Get returns the value of the variable.
//
// Returns:
//   - T: The value. The zero value if the receiver is nil.
func (s *Var[T]) Get() T {
	if s == nil {
		return *new(T)
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.data
}

// Set sets the value of the variable.
//
// Parameters:
//   - data: The value to set.
//
// Returns:
//   - error: An error if the receiver is nil.
func (s *Var[T]) Set(data T) error {
	if s == nil {
		return common.ErrNilReceiver
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = data

	return nil
}

// Edit edits the value of the variable.
//
// Parameters:
//   - fn: The function to edit the value.
//
// Returns:
//   - error: An error if the receiver is nil or the function is nil.
func (s *Var[T]) Edit(fn func(elem *T)) error {
	if fn == nil {
		return nil
	} else if s == nil {
		return common.ErrNilReceiver
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	fn(&s.data)

	return nil
}
