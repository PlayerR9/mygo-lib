package sets

import (
	"cmp"
	"iter"
	"slices"

	"github.com/PlayerR9/mygo-lib/OLD/common"
)

// OrderedSet is an ordered set in ascending order. An empty ordered set can
// either be created with the `var set OrderedSet[T]` syntax or with the
// `new(OrderedSet[T])` constructor.
type OrderedSet[T cmp.Ordered] struct {
	elems    []T
	lenElems uint
}

// Size implements the Set interface.
func (s OrderedSet[T]) Size() uint {
	return s.lenElems
}

// IsEmpty implements the Set interface.
func (s OrderedSet[T]) IsEmpty() bool {
	return s.lenElems == 0
}

// Reset implements the Set interface.
func (s *OrderedSet[T]) Reset() {
	if s == nil || s.lenElems == 0 {
		return
	}

	clear(s.elems)
	s.elems = nil
	s.lenElems = 0
}

// Add implements the Set interface.
func (s *OrderedSet[T]) Add(elem T) error {
	if s == nil {
		return common.ErrNilReceiver
	}

	pos, ok := slices.BinarySearch(s.elems, elem)
	if ok {
		return nil
	}

	s.elems = slices.Insert(s.elems, pos, elem)
	s.lenElems++

	return nil
}

// Contains implements the Set interface.
func (s OrderedSet[T]) Contains(elem T) bool {
	if s.lenElems == 0 {
		return false
	}

	_, ok := slices.BinarySearch(s.elems, elem)
	return ok
}

// Elem implements the Set interface.
func (s OrderedSet[T]) Elem() iter.Seq[T] {
	if len(s.elems) == 0 {
		return func(yield func(T) bool) {}
	}

	return func(yield func(T) bool) {
		for _, elem := range s.elems {
			if !yield(elem) {
				return
			}
		}
	}
}

// NewOrderedSet creates a new ordered set from the provided elements.
// The set will contain unique elements in ascending order.
//
// Parameters:
//   - elems: A slice of elements to initialize the set. Elements must be of an ordered type.
//
// Returns:
//   - *OrderedSet[T]: A pointer to the newly created ordered set containing unique elements.
func NewOrderedSet[T cmp.Ordered](elems []T) *OrderedSet[T] {
	if len(elems) < 2 {
		return &OrderedSet[T]{
			elems: elems,
		}
	}

	unique := make([]T, 0, len(elems))

	for _, elem := range elems {
		pos, ok := slices.BinarySearch(unique, elem)
		if ok {
			continue
		}

		unique = slices.Insert(unique, pos, elem)
	}

	lenElems := uint(len(unique))

	return &OrderedSet[T]{
		elems:    unique[:lenElems:lenElems],
		lenElems: lenElems,
	}
}
