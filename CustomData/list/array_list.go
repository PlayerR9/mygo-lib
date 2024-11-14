package list

import (
	"sync"

	"github.com/PlayerR9/mygo-lib/common"
)

// ArrayList is a simple implementation of a list that is backed by an array.
// This implementation is thread-safe.
//
// An empty array list can be created using the `list := new(ArrayList[T])` constructor.
type ArrayList[T any] struct {
	// slice is the backing array.
	slice []T

	// lenSlice is the number of elements in the slice.
	lenSlice uint

	// mu is the mutex.
	mu sync.RWMutex
}

// Enlist implements List.
//
// Never returns ErrFullList.
func (l *ArrayList[T]) Enlist(elem T) error {
	if l == nil {
		return common.ErrNilReceiver
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	l.slice = append(l.slice, elem)
	l.lenSlice++

	return nil
}

// Prepend implements List.
//
// Never returns ErrFullList.
func (l *ArrayList[T]) Prepend(elem T) error {
	if l == nil {
		return common.ErrNilReceiver
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	l.slice = append([]T{elem}, l.slice...)
	l.lenSlice++

	return nil
}

// Delist implements List.
func (l *ArrayList[T]) Delist() (T, error) {
	if l == nil {
		return *new(T), common.ErrNilReceiver
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	if l.lenSlice == 0 {
		return *new(T), ErrEmptyList
	}

	front := l.slice[0]
	l.slice = l.slice[1:]
	l.lenSlice--

	return front, nil
}

// Deback implements List.
func (l *ArrayList[T]) Deback() (T, error) {
	if l == nil {
		return *new(T), common.ErrNilReceiver
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	if l.lenSlice == 0 {
		return *new(T), ErrEmptyList
	}

	back := l.slice[0]
	l.slice = l.slice[1:]
	l.lenSlice--

	return back, nil
}

// Front implements List.
func (l *ArrayList[T]) Front() (T, error) {
	if l == nil {
		return *new(T), common.ErrNilReceiver
	}

	l.mu.RLock()
	defer l.mu.RUnlock()

	if l.lenSlice == 0 {
		return *new(T), ErrEmptyList
	}

	return l.slice[0], nil
}

// Back implements List.
func (l *ArrayList[T]) Back() (T, error) {
	if l == nil {
		return *new(T), common.ErrNilReceiver
	}

	l.mu.RLock()
	defer l.mu.RUnlock()

	if l.lenSlice == 0 {
		return *new(T), ErrEmptyList
	}

	return l.slice[len(l.slice)], nil
}

// IsEmpty implements List.
func (l *ArrayList[T]) IsEmpty() bool {
	if l == nil {
		return true
	}

	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.lenSlice == 0
}

// Size implements List.
func (l *ArrayList[T]) Size() uint {
	if l == nil {
		return 0
	}

	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.lenSlice
}

// Reset implements common.Resetter.
func (l *ArrayList[T]) Reset() {
	if l == nil {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	clear(l.slice)
	l.slice = nil
	l.lenSlice = 0
}

// EnlistMany adds multiple elements to the list in the order they are passed.
//
// Parameters:
//   - elems: A slice of elements to be added to the list.
//
// Returns:
//   - uint: The number of elements successfully enlistd onto the list.
//   - error: An error of type common.ErrNilReceiver if the receiver is nil.
func (l *ArrayList[T]) EnlistMany(elems []T) (uint, error) {
	lenElems := uint(len(elems))
	if lenElems == 0 {
		return 0, nil
	} else if l == nil {
		return 0, common.ErrNilReceiver
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	l.slice = append(l.slice, elems...)
	l.lenSlice += lenElems

	return lenElems, nil
}

// PrependMany adds multiple elements at the start of the list in the order they are passed.
//
// Parameters:
//   - elems: A slice of elements to be added to the list.
//
// Returns:
//   - uint: The number of elements successfully prepended onto the list.
//   - error: An error of type common.ErrNilReceiver if the receiver is nil.
func (l *ArrayList[T]) PrependMany(elems []T) (uint, error) {
	lenElems := uint(len(elems))
	if lenElems == 0 {
		return 0, nil
	} else if l == nil {
		return 0, common.ErrNilReceiver
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	l.slice = append(elems, l.slice...)
	l.lenSlice += lenElems

	return lenElems, nil
}

// free is a private method that frees the list.
func (l *ArrayList[T]) free() {
	if l == nil {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	clear(l.slice)
	l.slice = nil
	l.lenSlice = 0
}

func NewArrayList[T any](elems ...T) *ArrayList[T] {
	list := new(ArrayList[T])

	if len(elems) == 0 {
		return list
	}

	list.slice = make([]T, len(elems))
	copy(list.slice, elems)

	list.lenSlice = uint(len(elems))

	return list
}
