package slices

import "github.com/PlayerR9/mygo-lib/common"

// Predicate is a function that returns true if the element is valid.
//
// Parameters:
//   - e: The element to check.
//
// Returns:
//   - bool: True if the element is valid, false otherwise.
type Predicate[E any] func(e E) bool

// Filter filters a slice of E's according to a predicate.
//
// Parameters:
//   - s: The slice of E's to filter.
//   - predicate: The predicate to use for filtering.
//
// Returns:
//   - error: An error if the predicate is nil.
//
// Errors:
//   - common.ErrBadParam: If the receiver is nil.
func Filter[S ~[]E, E any](s *S, predicate Predicate[E]) error {
	if predicate == nil {
		err := common.NewErrNilParam("predicate")
		return err
	}

	if s == nil || len(*s) == 0 {
		return nil
	}

	var end uint

	for _, v := range *s {
		ok := predicate(v)
		if ok {
			(*s)[end] = v
			end++
		}
	}

	if end == 0 {
		clear(*s)
		*s = nil
	} else {
		clear((*s)[end:])
		*s = (*s)[:end]
	}

	return nil
}

// Reject rejects a slice of E's according to a predicate.
//
// Parameters:
//   - s: The slice of E's to reject.
//   - predicate: The predicate to use for rejecting.
//
// Returns:
//   - error: An error if the predicate is nil.
//
// Errors:
//   - common.ErrBadParam: If the receiver is nil.
func Reject[S ~[]E, E any](s *S, predicate Predicate[E]) error {
	if predicate == nil {
		err := common.NewErrNilParam("predicate")
		return err
	}

	if s == nil || len(*s) == 0 {
		return nil
	}

	var end uint

	for _, v := range *s {
		ok := predicate(v)
		if !ok {
			(*s)[end] = v
			end++
		}
	}

	if end == 0 {
		clear(*s)
		*s = nil
	} else {
		clear((*s)[end:])
		*s = (*s)[:end]
	}

	return nil
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

	var end uint

	for _, v := range *s {
		if v != nil {
			(*s)[end] = v
			end++
		}
	}

	n := uint(len(*s)) - end

	if end == 0 {
		clear(*s)
		*s = nil
	} else {
		clear((*s)[end:])
		*s = (*s)[:end]
	}

	return n
}

// RejectEmpty rejects in-place all zero-valued elements from a slice of E's and returns the number
// of elements rejected.
//
// Parameters:
//   - s: The slice of E's to filter.
//
// Returns:
//   - uint: The number of elements rejected.
func RejectEmpty[S ~[]E, E comparable](s *S) uint {
	if s == nil || len(*s) == 0 {
		return 0
	}

	zero := *new(E)

	var end uint

	for _, e := range *s {
		if e != zero {
			(*s)[end] = e
			end++
		}
	}

	n := uint(len(*s)) - end

	if end == 0 {
		clear(*s)
		*s = nil
	} else {
		clear((*s)[end:])
		*s = (*s)[:end]
	}

	return n
}
