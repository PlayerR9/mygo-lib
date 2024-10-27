package slices

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
// Behavior:
//   - If the list is empty, the predicate is nil, or there is no element
//     that satisfies the predicate, the slice is cleared and set to nil.
func Filter[T any](slice *[]T, p Predicate[T]) {
	if slice == nil || len(*slice) == 0 {
		return
	} else if p == nil {
		clear(*slice)
		*slice = nil

		return
	}

	var top int

	for _, elem := range *slice {
		ok := p(elem)
		if ok {
			(*slice)[top] = elem
			top++
		}
	}

	if top == 0 {
		clear(*slice)
		*slice = nil

		return
	}

	clear((*slice)[top:])
	*slice = (*slice)[:top:top]
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
func FilterIfApplicable[T any](slice *[]T, p Predicate[T]) bool {
	if slice == nil || len(*slice) == 0 {
		return true
	} else if p == nil {
		return false
	}

	var top int

	for _, elem := range *slice {
		ok := p(elem)
		if ok {
			(*slice)[top] = elem
			top++
		}
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
//   - []T: the list of elements that do not satisfy the predicate.
//
// Behavior:
//   - If the list is empty or the predicate is nil, returns nil.
func Reject[T any](slice *[]T, p Predicate[T]) {
	if slice == nil || len(*slice) == 0 {
		return
	} else if p == nil {
		clear(*slice)
		*slice = nil

		return
	}

	var top int

	for _, elem := range *slice {
		ok := p(elem)
		if !ok {
			(*slice)[top] = elem
			top++
		}
	}

	if top == 0 {
		clear(*slice)
		*slice = nil

		return
	}

	clear((*slice)[top:])
	*slice = (*slice)[:top:top]
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

	var top int

	for _, elem := range *slice {
		ok := p(elem)
		if !ok {
			(*slice)[top] = elem
			top++
		}
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
func RejectNils[T any](slice *[]*T) {
	if slice == nil || len(*slice) == 0 {
		return
	}

	var top int

	for _, elem := range *slice {
		if elem != nil {
			(*slice)[top] = elem
			top++
		}
	}

	if top == 0 {
		clear(*slice)
		*slice = nil

		return
	}

	clear((*slice)[top:])
	*slice = (*slice)[:top:top]
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
// Behavior:
//   - If the provided slice is empty or the filter function is nil, the original slice is cleared and set to nil.
//   - The filter function is called repeatedly with the current list of indices until it returns true or the list of indices is empty.
//   - The filtered slice contains only the elements corresponding to the selected indices.
func ComplexFilter[T any](slice *[]T, fn func(indices *[]int) bool) {
	if slice == nil || len(*slice) == 0 {
		return
	} else if fn == nil {
		clear(*slice)
		*slice = nil

		return
	}

	indices := make([]int, 0, len(*slice))
	for i := range *slice {
		indices = append(indices, i)
	}

	var early_exit bool

	for len(indices) > 0 && !early_exit {
		early_exit = fn(&indices)
	}

	if !early_exit {
		clear(*slice)
		*slice = nil

		return
	}

	var top int

	for _, idx := range indices {
		(*slice)[top] = (*slice)[idx]
		top++
	}

	clear((*slice)[top:])
	*slice = (*slice)[:top:top]
}
