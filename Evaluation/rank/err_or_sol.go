package rank

import (
	"iter"

	"github.com/PlayerR9/mygo-lib/common"
)

// ErrOrSol is an ordered set according to rank.
type ErrOrSol[T any] struct {
	// errs is the list of errors.
	errs *Rank[error]

	// sols is the list of solutions.
	sols *Rank[T]
}

func NewErrOrSol[T any]() *ErrOrSol[T] {
	return &ErrOrSol[T]{
		errs: new(Rank[error]),
		sols: nil,
	}
}

// Size returns the number of elements in the set.
//
// Returns:
//   - int: The number of elements in the set. Never negative.
//
// If the receiver is nil, 0 is returned.
func (s *ErrOrSol[T]) Size() int {
	if s == nil {
		return 0
	}

	if s.errs == nil {
		return s.sols.Size()
	} else {
		return s.errs.Size()
	}
}

// IsEmpty returns whether the set is empty.
//
// Returns:
//   - bool: True if the set is empty, false otherwise.
//
// If the receiver is nil, true is returned.
func (s *ErrOrSol[T]) IsEmpty() bool {
	if s == nil {
		return true
	}

	if s.errs == nil {
		return s.sols.IsEmpty()
	} else {
		return s.errs.IsEmpty()
	}
}

// Reset resets the set for reuse.
func (s *ErrOrSol[T]) Reset() {
	if s == nil {
		return
	}

	if s.errs != nil {
		s.errs.Reset()
		s.errs = new(Rank[error])
	}

	if s.sols != nil {
		s.sols.Reset()
		s.sols = nil
	}
}

// AddSol adds an element to the set.
//
// Parameters:
//   - rank: The rank of the element.
//   - elem: The element to add.
//
// Returns:
//   - error: An error if the receiver is nil.
func (s *ErrOrSol[T]) AddSol(rank int, elem T) error {
	if s == nil {
		return common.ErrNilReceiver
	}

	if s.errs != nil {
		s.errs.Reset()
		s.errs = nil
	}

	if s.sols == nil {
		s.sols = new(Rank[T])
	}

	_ = s.sols.Add(rank, elem)

	return nil
}

// AddErr adds an element to the set.
//
// Parameters:
//   - rank: The rank of the element.
//   - elem: The element to add.
//
// Returns:
//   - error: An error if the receiver is nil.
func (s *ErrOrSol[T]) AddErr(rank int, err error) error {
	if err == nil {
		return nil
	} else if s == nil {
		return common.ErrNilReceiver
	}

	if s.errs == nil {
		return nil
	}

	_ = s.errs.Add(rank, err)

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
func (s *ErrOrSol[T]) Elem() iter.Seq2[int, T] {
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
func (s *ErrOrSol[T]) ReverseElem() iter.Seq2[int, T] {
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
