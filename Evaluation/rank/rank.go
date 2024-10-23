package rank

import (
	"iter"
	"slices"

	"github.com/PlayerR9/mygo-lib/common"
)

// Rank is an ordered set according to rank. An empty rank set can
// either be created with the `var set Rank[T]` syntax or with the
// `new(Rank[T])` constructor.
type Rank[T any] struct {
	// buckets maps ranks to buckets.
	buckets map[int][]T

	// ranks is a sorted list of ranks.
	ranks []int

	// size is the number of elements in the set.
	size int
}

// Size returns the number of elements in the set.
//
// Returns:
//   - int: The number of elements in the set. Never negative.
//
// If the receiver is nil, 0 is returned.
func (s *Rank[T]) Size() int {
	if s == nil {
		return 0
	}

	return s.size
}

// IsEmpty returns whether the set is empty.
//
// Returns:
//   - bool: True if the set is empty, false otherwise.
//
// If the receiver is nil, true is returned.
func (s *Rank[T]) IsEmpty() bool {
	return s == nil || s.size == 0
}

// Reset resets the set for reuse.
func (s *Rank[T]) Reset() {
	if s == nil {
		return
	}

	if len(s.buckets) > 0 {
		for k, v := range s.buckets {
			clear(v)
			s.buckets[k] = nil
		}

		clear(s.buckets)
		s.buckets = nil
	}

	if len(s.ranks) > 0 {
		clear(s.ranks)
		s.ranks = nil
	}

	s.size = 0
}

// Add adds an element to the set.
//
// Parameters:
//   - rank: The rank of the element.
//   - elem: The element to add.
//
// Returns:
//   - error: An error if the receiver is nil.
func (s *Rank[T]) Add(rank int, elem T) error {
	if s == nil {
		return common.ErrNilReceiver
	}

	if s.buckets == nil {
		s.buckets = make(map[int][]T)
	}

	bucket, ok := s.buckets[rank]
	if !ok {
		s.buckets[rank] = []T{elem}

		pos, ok := slices.BinarySearch(s.ranks, rank)
		if !ok {
			s.ranks = slices.Insert(s.ranks, pos, rank)
		}
	} else {
		s.buckets[rank] = append(bucket, elem)
	}

	s.size++

	return nil
}

// Elem iterates over the elements in the set in rank order, starting from the
// highest rank. The order of elements within a rank is guaranteed to be
// the same as the order in which they were added.
//
// Returns:
//   - iter.Seq2[int, T]: The elements in the set. Never returns nil.
//
// If the receiver is nil or the set is empty, an empty sequence is returned.
func (s *Rank[T]) Elem() iter.Seq2[int, T] {
	if s == nil || s.size == 0 {
		return func(yield func(int, T) bool) {}
	}

	return func(yield func(int, T) bool) {
		for i := len(s.ranks) - 1; i >= 0; i-- {
			rank := s.ranks[i]

			bucket := s.buckets[rank]

			for _, elem := range bucket {
				if !yield(rank, elem) {
					return
				}
			}
		}
	}
}

// ReverseElem iterates over the elements in the set in reverse rank order,
// starting from the lowest rank. The order of elements within a rank is
// guaranteed to be the same as the order in which they were added.
//
// Returns:
//   - iter.Seq2[int, T]: The elements in the set. Never returns nil.
//
// If the receiver is nil or the set is empty, an empty sequence is returned.
func (s *Rank[T]) ReverseElem() iter.Seq2[int, T] {
	if s == nil || s.size == 0 {
		return func(yield func(int, T) bool) {}
	}

	return func(yield func(int, T) bool) {
		for _, rank := range s.ranks {
			bucket := s.buckets[rank]

			for i := len(bucket) - 1; i >= 0; i-- {
				if !yield(rank, bucket[i]) {
					return
				}
			}
		}
	}
}

func (s *Rank[T]) Build() []T {
	if s == nil || s.size == 0 {
		return nil
	}

	slice := make([]T, 0, s.size)

	for i := len(s.ranks) - 1; i >= 0; i-- {
		rank := s.ranks[i]

		bucket := s.buckets[rank]

		for _, elem := range bucket {
			slice = append(slice, elem)
		}
	}

	return slice
}

// ReverseElem iterates over the elements in the set in reverse rank order,
// starting from the lowest rank. The order of elements within a rank is
// guaranteed to be the same as the order in which they were added.
//
// Returns:
//   - iter.Seq2[int, T]: The elements in the set. Never returns nil.
//
// If the receiver is nil or the set is empty, an empty sequence is returned.
func (s *Rank[T]) ReverseBuild() []T {
	if s == nil || s.size == 0 {
		return nil
	}

	slice := make([]T, 0, s.size)

	for _, rank := range s.ranks {
		bucket := s.buckets[rank]

		for i := len(bucket) - 1; i >= 0; i-- {
			slice = append(slice)

			if !yield(rank, bucket[i]) {
				return
			}
		}
	}
}
