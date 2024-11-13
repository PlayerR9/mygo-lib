package sets

import (
	"sync"

	"github.com/PlayerR9/mygo-lib/common"
)

// SeenSet is a set of elements that have been seen. An empty set can be created
// using the `set := new(SeenSet[T])` constructor.
//
// This set is thread-safe.
type SeenSet[T comparable] struct {
	// seen is the underlying set of seen elements.
	seen map[T]struct{}

	// mu is the mutex for the seen set.
	mu sync.RWMutex
}

// Add implements Set.
func (s *SeenSet[T]) Add(elem T) error {
	if s == nil {
		return common.ErrNilReceiver
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.seen == nil {
		s.seen = make(map[T]struct{})
	}

	s.seen[elem] = struct{}{}

	return nil
}

// Contains implements Set.
func (s *SeenSet[T]) Contains(elem T) bool {
	if s == nil {
		return false
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	if len(s.seen) == 0 {
		return false
	}

	_, ok := s.seen[elem]
	return ok
}

// See adds an element to the set. However, if the element is already in the set,
// it returns false.
//
// Parameters:
//   - elem: The element to add to the set.
//
// Returns:
//   - bool: True if the element was added to the set, false otherwise.
//   - error: Returns an error of type common.ErrNilReceiver if the receiver is nil.
func (s *SeenSet[T]) See(elem T) (bool, error) {
	if s == nil {
		return false, common.ErrNilReceiver
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.seen == nil {
		s.seen = make(map[T]struct{})
	} else {
		_, ok := s.seen[elem]
		if ok {
			return false, nil
		}
	}

	s.seen[elem] = struct{}{}

	return true, nil
}
