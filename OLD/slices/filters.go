package slices

import "github.com/PlayerR9/mygo-lib/OLD/common"

// Predicate is a type of function that checks whether an element
// satisfies a given condition.
//
// Parameters:
//   - elem: the element to check.
//
// Returns:
//   - bool: True if the element satisfies the condition, false otherwise.
type Predicate[T any] func(elem T) bool

// Filter applies a predicate function on a slice of elements;
// keeping only those elements that satisfy the predicate. This
// function modifies the original list in-place.
//
// Parameters:
//   - slice: the list of elements to filter.
//   - p: the predicate function to apply.
//
// Returns:
//   - uint: the number of elements removed from the list.
//
// Behavior:
//   - If the list is empty, the predicate is nil, or there is no element
//     that satisfies the predicate, the slice is cleared and set to nil.
//
// Panics:
//   - common.ErrBadParam: If the slice is nil.
func Filter[T any](slice *[]T, p Predicate[T]) uint {
	if slice == nil {
		panic(common.NewErrNilParam("slice"))
	}

	lenSlice := uint(len(*slice))
	if lenSlice == 0 {
		return 0
	} else if p == nil {
		clear(*slice)
		*slice = nil

		return lenSlice
	}

	var top uint

	for _, elem := range *slice {
		ok := p(elem)
		if !ok {
			continue
		}

		(*slice)[top] = elem
		top++
	}

	if top == 0 {
		clear(*slice)
		*slice = nil

		return lenSlice
	}

	clear((*slice)[top:])
	*slice = (*slice)[:top:top]

	return lenSlice - top
}

// FilterIfApplicable applies a predicate function on a slice of elements, retaining only
// those elements that satisfy the predicate. This function modifies the
// original list in-place. Does nothing if no elements satisfy the predicate.
//
// Parameters:
//   - slice: The list of elements to filter.
//   - p: The predicate function to apply.
//
// Returns:
//   - bool: True if the slice was empty or elements were successfully
//     filtered, false if the predicate is nil or no elements satisfy the predicate.
//
// Behavior:
//   - If the list is empty or all elements are removed, returns true.
//   - If the predicate is nil or no elements satisfy the predicate, returns false.
//
// Panics:
//   - common.ErrBadParam: If the slice is nil.
func FilterIfApplicable[T any](slice *[]T, p Predicate[T]) bool {
	if slice == nil {
		panic(common.NewErrNilParam("slice"))
	}

	lenSlice := uint(len(*slice))
	if lenSlice == 0 {
		return true
	} else if p == nil {
		return false
	}

	var top uint

	for _, elem := range *slice {
		ok := p(elem)
		if !ok {
			continue
		}

		(*slice)[top] = elem
		top++
	}

	if top == 0 {
		return false
	}

	clear((*slice)[top:])
	*slice = (*slice)[:top:top]

	return true
}

// Reject applies a predicate function on a slice of elements;
// keeping only those elements that do not satisfy the predicate. This
// function modifies the original list in-place.
//
// Parameters:
//   - slice: the list of elements to filter.
//   - p: the predicate function to apply.
//
// Returns:
//   - uint: the number of elements removed from the list.
//
// Behavior:
//   - If the list is empty or the predicate is nil, returns nil.
//
// Panics:
//   - common.ErrBadParam: If the slice is nil.
func Reject[T any](slice *[]T, p Predicate[T]) uint {
	if slice == nil {
		panic(common.NewErrNilParam("slice"))
	}

	lenSlice := uint(len(*slice))
	if lenSlice == 0 {
		return 0
	} else if p == nil {
		clear(*slice)
		*slice = nil

		return lenSlice
	}

	var top uint

	for _, elem := range *slice {
		ok := p(elem)
		if ok {
			continue
		}

		(*slice)[top] = elem
		top++
	}

	if top == 0 {
		clear(*slice)
		*slice = nil

		return lenSlice
	}

	clear((*slice)[top:])
	*slice = (*slice)[:top:top]

	return lenSlice - top
}

// RejectIfApplicable applies a predicate function on a slice of elements, rejecting
// only those elements that do not satisfy the predicate. This function modifies the
// original list in-place. Does nothing if no elements do not satisfy the predicate.
//
// Parameters:
//   - slice: the list of elements to filter.
//   - p: the predicate function to apply.
//
// Returns:
//   - bool: True if the slice was empty or all elements were successfully
//     filtered, false if the predicate is nil or no elements do not satisfy the
//     predicate.
func RejectIfApplicable[T any](slice *[]T, p Predicate[T]) bool {
	if slice == nil || len(*slice) == 0 {
		return true
	} else if p == nil {
		return false
	}

	var top uint

	for _, elem := range *slice {
		ok := p(elem)
		if ok {
			continue
		}

		(*slice)[top] = elem
		top++
	}

	if top == 0 {
		return false
	}

	clear((*slice)[top:])
	*slice = (*slice)[:top:top]

	return true
}

// RejectNils works like Reject but keeps only non-nil elements.
//
// Parameters:
//   - slice: the list of elements to filter.
//
// Returns:
//   - uint: the number of elements removed from the list.
//
// Panics:
//   - common.ErrBadParam: If the slice is nil.
func RejectNils[T any](slice *[]*T) uint {
	if slice == nil {
		panic(common.NewErrNilParam("slice"))
	}

	lenSlice := uint(len(*slice))
	if lenSlice == 0 {
		return 0
	}

	var top uint

	for _, elem := range *slice {
		if elem == nil {
			continue
		}

		(*slice)[top] = elem
		top++
	}

	if top == 0 {
		clear(*slice)
		*slice = nil

		return lenSlice
	}

	clear((*slice)[top:])
	*slice = (*slice)[:top:top]

	return lenSlice - top
}

// ComplexFilter applies a filter function on a slice of elements based on the provided filter function.
// As with Filter, this function modifies the original list in-place.
//
// This function uses indices for optimization reasons.
//
// Parameters:
//   - slice: The slice of elements to filter.
//   - fn: The filter function that takes a list of indices and returns a boolean indicating whether to early exit.
//
// Returns:
//   - uint: The number of elements removed from the slice.
//
// Behavior:
//   - If the provided slice is empty or the filter function is nil, the original slice is cleared and set to nil.
//   - The filter function is called repeatedly with the current list of indices until it returns true or the list of indices is empty.
//   - The filtered slice contains only the elements corresponding to the selected indices.
//
// Panics:
//   - common.ErrBadParam: If the slice is nil.
func ComplexFilter[T any](slice *[]T, fn func(indices *[]uint) bool) uint {
	if slice == nil {
		panic(common.NewErrNilParam("slice"))
	}

	lenSlice := uint(len(*slice))
	if lenSlice == 0 {
		return 0
	} else if fn == nil {
		clear(*slice)
		*slice = nil

		return lenSlice
	}

	indices := make([]uint, 0, lenSlice)
	for i := uint(0); i < lenSlice; i++ {
		indices = append(indices, i)
	}

	var early_exit bool

	for len(indices) > 0 && !early_exit {
		early_exit = fn(&indices)
	}

	if !early_exit {
		clear(*slice)
		*slice = nil

		return lenSlice
	}

	var top uint

	for _, idx := range indices {
		(*slice)[top] = (*slice)[idx]
		top++
	}

	clear((*slice)[top:])
	*slice = (*slice)[:top:top]

	return lenSlice - top
}

// Split splits a slice into two parts based on a given predicate.
// The first part of the slice will contain all elements that satisfy the
// predicate, and the second part will contain all elements that do not.
//
// Parameters:
//   - slice: the elements to split.
//   - predicate: the predicate function to apply.
//
// Returns:
//   - uint: the number of elements in the first part of the slice.
func Split[T any](slice []T, predicate Predicate[T]) uint {
	lenSlice := uint(len(slice))
	if lenSlice == 0 {
		return 0
	} else if predicate == nil {
		return lenSlice
	}

	var bound uint

	for i := uint(0); i < lenSlice; i++ {
		elem := slice[i]

		ok := predicate(elem)
		if !ok {
			continue
		}

		for j := i; j > bound; j-- {
			slice[j] = slice[j-1]
		}

		slice[bound] = elem
		bound++
	}

	return bound
}
