package rws

import (
	"sync"

	"github.com/PlayerR9/mygo-lib/common"
)

// Slice is a slice that can be read and written safely.
type Slice[T any] struct {
	// slice is the slice to be read and written safely.
	slice []T

	// mu is the read-write lock for the slice.
	mu sync.RWMutex
}

// NewSlice creates a new Slice.
//
// Returns:
//   - *Slice[T]: The new Slice. Never returns nil.
func NewSlice[T any]() *Slice[T] {
	return &Slice[T]{}
}

// Append appends a slice to the Slice.
//
// Returns:
//   - error: Nil if the receiver is nil.
//
// Errors:
//   - errors.ErrNilReceiver: If the receiver is nil.
func (s *Slice[T]) Append(slice []T) error {
	if len(slice) == 0 {
		return nil
	} else if s == nil {
		return common.ErrNilReceiver
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.slice = append(s.slice, slice...)

	return nil
}

// Free releases any resources associated with the Slice.
func (s *Slice[T]) Free() {
	if s == nil {
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.slice) == 0 {
		return
	}

	clear(s.slice)
	s.slice = nil
}

// Size returns the number of elements in the Slice.
//
// Returns:
//   - int: The number of elements in the Slice. 0 when the receiver is nil.
func (s *Slice[T]) Size() int {
	if s == nil {
		return 0
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.slice)
}

// Slice returns a copy of the elements in the Slice.
//
// Returns:
//   - []T: A copy of the elements in the Slice.
func (s *Slice[T]) Slice() []T {
	if s == nil {
		return nil
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	slice := make([]T, len(s.slice))
	copy(slice, s.slice)

	return slice
}
