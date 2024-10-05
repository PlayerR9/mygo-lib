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

// ApplyFilter applies a predicate function on a list of elements;
// keeping only those elements that satisfy the predicate.
//
// Parameters:
//   - elems: the list of elements to filter.
//   - f: the predicate function to apply.
//
// Returns:
//   - []T: the list of elements that satisfy the predicate.
//
// Behaviors:
//   - If the list is empty or the predicate is nil, returns nil.
//
// WARNING: This function modifies the original list if there is at least
// one element that does not satisfy the predicate. Make sure to copy the
// list before calling this function if you want to keep the original list.
func ApplyFilter[T any](elems []T, f Predicate[T]) []T {
	if len(elems) == 0 || f == nil {
		return nil
	}

	var top int

	for i := 0; i < len(elems); i++ {
		if f(elems[i]) {
			elems[top] = elems[i]
			top++
		}
	}

	return elems[:top:top]
}

// ApplyReject applies a predicate function on a list of elements;
// keeping only those elements that do not satisfy the predicate.
//
// Parameters:
//   - elems: the list of elements to filter.
//   - f: the predicate function to apply.
//
// Returns:
//   - []T: the list of elements that do not satisfy the predicate.
//
// Behaviors:
//   - If the list is empty or the predicate is nil, returns nil.
//
// WARNING: This function modifies the original list if there is at least
// one element that satisfies the predicate. Make sure to copy the
// list before calling this function if you want to keep the original list.
func ApplyReject[T any](elems []T, f Predicate[T]) []T {
	if len(elems) == 0 || f == nil {
		return nil
	}

	var top int

	for i := 0; i < len(elems); i++ {
		if !f(elems[i]) {
			elems[top] = elems[i]
			top++
		}
	}

	return elems[:top:top]
}
