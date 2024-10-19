package sets

import (
	"cmp"
	"fmt"
	"iter"
	"slices"
	"strings"

	gers "github.com/PlayerR9/mygo-lib/errors"
)

// OrderedSet is an ordered set in ascending order.
type OrderedSet[K cmp.Ordered] struct {
	// elems are the elements in the set.
	elems []K
}

// String implements the Set interface.
func (s OrderedSet[K]) String() string {
	elems := make([]string, 0, len(s.elems))

	for _, elem := range s.elems {
		elems = append(elems, fmt.Sprint(elem))
	}

	return "OrderedSet[" + strings.Join(elems, ", ") + "]"
}

// Size implements the Set interface.
func (s OrderedSet[K]) Size() int {
	return len(s.elems)
}

// IsEmpty implements the Set interface.
func (s OrderedSet[K]) IsEmpty() bool {
	return len(s.elems) == 0
}

// NewOrderedSet creates a new ordered set.
//
// Returns:
//   - *OrderedSet[K]: The new ordered set. Never returns nil.
//
// This function is just for convenience since it does the same as:
//
//	var set OrderedSet[K]
func NewOrderedSet[K cmp.Ordered]() *OrderedSet[K] {
	return &OrderedSet[K]{}
}

// NewOrderedSetFromSlice creates a new ordered set from a slice.
//
// Parameters:
//   - slice: The elements to add to the set.
//
// Returns:
//   - *OrderedSet[K]: The new ordered set. Never returns nil.
func NewOrderedSetFromSlice[K cmp.Ordered](slice []K) *OrderedSet[K] {
	elems := make([]K, 0, len(slice))

	for _, elem := range slice {
		pos, ok := slices.BinarySearch(elems, elem)
		if !ok {
			elems = slices.Insert(elems, pos, elem)
		}
	}

	return &OrderedSet[K]{
		elems: elems,
	}
}

// Add adds an element to the set.
//
// Parameters:
//   - k: The element to add.
//
// Returns:
//   - error: An error if the element could not be added to the set.
//
// Errors:
//   - errors.ErrNilReceiver: If the receiver is nil.
func (s *OrderedSet[K]) Add(k K) error {
	if s == nil {
		return gers.ErrNilReceiver
	}

	pos, ok := slices.BinarySearch(s.elems, k)
	if !ok {
		s.elems = slices.Insert(s.elems, pos, k)
	}

	return nil
}

// AddMany adds multiple elements to the set.
//
// Parameters:
//   - ks: The elements to add.
//
// Returns:
//   - error: An error if the element could not be added to the set.
//
// Errors:
//   - errors.ErrNilReceiver: If the receiver is nil.
func (s *OrderedSet[K]) AddMany(ks ...K) error {
	if len(ks) == 0 {
		return nil
	} else if s == nil {
		return gers.ErrNilReceiver
	}

	for _, k := range ks {
		pos, ok := slices.BinarySearch(s.elems, k)
		if !ok {
			s.elems = slices.Insert(s.elems, pos, k)
		}
	}

	return nil
}

// Elem iterates through the elements in the set.
//
// Returns:
//   - iter.Seq[K]: The elements in the set. Never returns nil.
func (s OrderedSet[K]) Elem() iter.Seq[K] {
	return func(yield func(K) bool) {
		for _, elem := range s.elems {
			if !yield(elem) {
				return
			}
		}
	}
}

// Has checks if an element is in the set.
//
// Parameters:
//   - k: The element to check.
//
// Returns:
//   - bool: True if the element is in the set, false otherwise.
func (s OrderedSet[K]) Has(k K) bool {
	_, ok := slices.BinarySearch(s.elems, k)
	return ok
}

// Merge merges another set into this set. Does nothing if the other set is empty or nil.
//
// Parameters:
//   - other: The other set to merge into this set.
//
// Returns:
//   - error: An error if the receiver is nil.
func (s *OrderedSet[K]) Merge(other *OrderedSet[K]) error {
	if other == nil || other.IsEmpty() {
		return nil
	} else if s == nil {
		return gers.ErrNilReceiver
	}

	for _, elem := range other.elems {
		pos, ok := slices.BinarySearch(s.elems, elem)
		if !ok {
			s.elems = slices.Insert(s.elems, pos, elem)
		}
	}

	return nil
}
