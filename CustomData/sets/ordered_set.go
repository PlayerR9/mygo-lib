package sets

import (
	"cmp"
	"fmt"
	"iter"
	"slices"
	"strings"

	gers "github.com/PlayerR9/mygo-lib/errors"
)

type OrderedSet[K cmp.Ordered] struct {
	elems []K
}

func (s OrderedSet[K]) String() string {
	elems := make([]string, 0, len(s.elems))

	for _, elem := range s.elems {
		elems = append(elems, fmt.Sprintf("%v", elem))
	}

	return "OrderedSet[" + strings.Join(elems, ", ") + "]"
}

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

func (s OrderedSet[K]) Size() int {
	return len(s.elems)
}

func (s OrderedSet[K]) Elem() iter.Seq[K] {
	return func(yield func(K) bool) {
		for _, elem := range s.elems {
			if !yield(elem) {
				return
			}
		}
	}
}

func (s OrderedSet[K]) Has(k K) bool {
	_, ok := slices.BinarySearch(s.elems, k)
	return ok
}

func (s OrderedSet[K]) IsEmpty() bool {
	return len(s.elems) == 0
}

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
