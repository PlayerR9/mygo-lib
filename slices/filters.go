package slices

import (
	"github.com/PlayerR9/mygo-lib/slices/internal"
)

// Filter filters a slice of E's according to a predicate.
//
// Parameters:
//   - s: The slice of E's to filter.
//   - predicate: The predicate to use for filtering.
//
// Returns:
//   - uint: The number of elements filtered-out.
//
// Panics:
//   - ErrNoPredicate: If the predicate is nil.
func Filter[S ~[]E, E any](s *S, predicate func(e E) bool) uint {
	if predicate == nil {
		panic(ErrNoPredicate)
	}

	if s == nil || len(*s) == 0 {
		return 0
	}

	n := internal.Filter(s, predicate)
	return n
}

// Reject rejects a slice of E's according to a predicate.
//
// Parameters:
//   - s: The slice of E's to reject.
//   - predicate: The predicate to use for rejecting.
//
// Returns:
//   - uint: The number of elements rejected.
//
// Panics:
//   - ErrNoPredicate: If the predicate is nil.
func Reject[S ~[]E, E any](s *S, predicate func(e E) bool) uint {
	if predicate == nil {
		panic(ErrNoPredicate)
	}

	if s == nil || len(*s) == 0 {
		return 0
	}

	n := internal.Reject(s, predicate)
	return n
}

// RejectNils rejects in-place all nil elements from a slice of E's and returns the number
// of elements rejected.
//
// Parameters:
//   - s: The slice of E's to filter.
//
// Returns:
//   - uint: The number of elements rejected.
func RejectNils[S ~[]*E, E any](s *S) uint {
	if s == nil || len(*s) == 0 {
		return 0
	}

	n := internal.RejectNils(s)
	return n
}

// RejectZero rejects in-place all zero-valued elements from a slice of E's and returns the number
// of elements rejected.
//
// Parameters:
//   - s: The slice of E's to filter.
//
// Returns:
//   - uint: The number of elements rejected.
func RejectZero[S ~[]E, E comparable](s *S) uint {
	if s == nil || len(*s) == 0 {
		return 0
	}

	n := internal.RejectZero(s)
	return n
}
