package sets

import "github.com/PlayerR9/mygo-lib/common"

// Set is an interface used by any set-like data structure.
type Set[T comparable] interface {
	// Add inserts an element into the set, if it is not already present. It
	// does not return an error if the element is already present.
	//
	// Parameters:
	//   - elem: The element to add to the set.
	//
	// Returns:
	//   - error: Returns an error if the element could not be added to the set.
	//
	// Errors:
	//   - common.ErrNilReceiver: If the receiver is nil.
	//   - any other error that may be returned.
	Add(elem T) error

	// Has checks whether the specified element is present in the set. If the
	// receiver is nil, it returns false.
	//
	// Parameters:
	//   - elem: The element to check for presence in the set.
	//
	// Returns:
	//   - bool: True if the element is present in the set, false otherwise.
	Has(elem T) bool
}

// Add adds multiple elements to the set, if they are not already present. It
// does not return an error if an element is already present.
//
// Parameters:
//   - s: The set to which the elements are added.
//   - elems: Variadic parameters representing the elements to be added.
//
// Returns:
//   - uint: The number of elements successfully added to the set.
//   - error: An error if one of the elements could not be added to the set.
//
// Errors:
//   - common.ErrBadParam: If the set is nil.
//   - any other error that may be returned by the `Add` method of the set.
func Add[T comparable](s Set[T], elems ...T) (uint, error) {
	if len(elems) == 0 {
		return 0, nil
	} else if s == nil {
		return 0, common.NewErrNilParam("s")
	}

	for i, elem := range elems {
		err := s.Add(elem)
		if err != nil {
			return uint(i), err
		}
	}

	return uint(len(elems)), nil
}

// Filter filters a slice of elements using the given set. All elements that
// are also present in the set are removed from the slice. The original slice
// is modified in-place.
//
// Parameters:
//   - s: The set to use for filtering.
//   - slice: The slice of elements to filter.
//
// Returns:
//   - uint: The number of elements removed from the slice.
//   - error: An error if the slice is nil.
//
// Behavior:
//   - If the set is empty or the slice is empty, does nothing and returns 0.
//   - If the slice is nil, returns an error.
//   - If the set is nil, returns an error.
//
// Errors:
//   - common.ErrBadParam: If the slice is nil or the set is nil.
func Filter[T comparable](s Set[T], slice *[]T) (uint, error) {
	if slice == nil {
		return 0, common.NewErrNilParam("slice")
	}

	lenSlice := uint(len(*slice))
	if lenSlice == 0 {
		return 0, nil
	} else if s == nil {
		return 0, common.NewErrNilParam("s")
	}

	var top uint

	for i := uint(0); i < lenSlice; i++ {
		elem := (*slice)[i]

		ok := s.Has(elem)
		if ok {
			continue
		}

		(*slice)[top] = elem
		top++
	}

	if top == 0 {
		clear(*slice)
		*slice = nil

		return lenSlice, nil
	} else {
		clear((*slice)[top:])
		*slice = (*slice)[:top]

		return lenSlice - top, nil
	}
}

// Reject filters a slice of elements using the given set. All elements that
// are not present in the set are removed from the slice. The original slice
// is modified in-place.
//
// Parameters:
//   - s: The set to use for filtering.
//   - slice: The slice of elements to filter.
//
// Returns:
//   - uint: The number of elements removed from the slice.
//   - error: An error if the slice is nil.
//
// Behavior:
//   - If the set is empty or the slice is empty, does nothing and returns 0.
//   - If the slice is nil, returns an error.
//   - If the set is nil, returns an error.
//
// Errors:
//   - common.ErrBadParam: If the slice is nil or the set is nil.
func Reject[T comparable](s Set[T], slice *[]T) (uint, error) {
	if slice == nil {
		return 0, common.NewErrNilParam("slice")
	}

	lenSlice := uint(len(*slice))
	if lenSlice == 0 {
		return 0, nil
	} else if s == nil {
		return 0, common.NewErrNilParam("s")
	}

	var top uint

	for i := uint(0); i < lenSlice; i++ {
		elem := (*slice)[i]

		ok := s.Has(elem)
		if !ok {
			continue
		}

		(*slice)[top] = elem
		top++
	}

	if top == 0 {
		clear(*slice)
		*slice = nil

		return lenSlice, nil
	} else {
		clear((*slice)[top:])
		*slice = (*slice)[:top]

		return lenSlice - top, nil
	}
}
