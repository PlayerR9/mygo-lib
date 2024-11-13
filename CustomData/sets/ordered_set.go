package sets

import (
	"cmp"
	"slices"
	"sync"

	"github.com/PlayerR9/mygo-lib/common"
)

// OrderedSet is an ordered set in ascending order. An empty ordered set can
// be created using the `set := new(OrderedSet[T])` constructor.
//
// This set is thread-safe.
type OrderedSet[T cmp.Ordered] struct {
	// elems is a slice of elements in ascending order.
	elems []T

	// mu is a mutex used to synchronize access to the set.
	mu sync.RWMutex
}

// Add implements Set.
func (s *OrderedSet[T]) Add(elem T) error {
	if s == nil {
		return common.ErrNilReceiver
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	pos, ok := slices.BinarySearch(s.elems, elem)
	if ok {
		return nil
	}

	s.elems = slices.Insert(s.elems, pos, elem)

	return nil
}

// Contains implements Set.
func (s *OrderedSet[T]) Contains(elem T) bool {
	if s == nil {
		return false
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	_, ok := slices.BinarySearch(s.elems, elem)
	return ok
}
