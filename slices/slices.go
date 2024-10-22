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
// keeping only those elements that satisfy the predicate.
//
// Parameters:
//   - elems: the list of elements to filter.
//   - f: the predicate function to apply.
//
// Returns:
//   - []T: the list of elements that satisfy the predicate.
//
// Behavior:
//   - If the list is empty, the predicate is nil, or there is no element
//     that satisfies the predicate, returns nil.
//   - WARNING: As a side-effect, the original list will be modified if there
//     is at least one element that satisfies the predicate and one that does
//     not.
func Filter[T any](slice []T, p Predicate[T]) []T {
	if len(slice) == 0 || p == nil {
		return nil
	}

	var top int

	for i := range slice {
		ok := p(slice[i])
		if ok {
			slice[top] = slice[i]
			top++
		}
	}

	if top == 0 {
		return nil
	}

	slice = slice[:top:top]

	return slice
}

// Reject applies a predicate function on a slice of elements;
// keeping only those elements that do not satisfy the predicate.
//
// Parameters:
//   - elems: the list of elements to filter.
//   - f: the predicate function to apply.
//
// Returns:
//   - []T: the list of elements that do not satisfy the predicate.
//
// Behavior:
//   - If the list is empty or the predicate is nil, returns nil.
//   - WARNING: As a side-effect, the original list will be modified if there
//     is at least one element that satisfies the predicate and one that does
//     not.
func Reject[T any](slice []T, p Predicate[T]) []T {
	if len(slice) == 0 || p == nil {
		return nil
	}

	var top int

	for i := range slice {
		ok := p(slice[i])
		if !ok {
			slice[top] = slice[i]
			top++
		}
	}

	if top == 0 {
		return nil
	}

	slice = slice[:top:top]

	return slice
}

// RejectNils works like Reject but returns nil if the list is empty. However,
// unlike Reject, it has no side-effects.
//
// Parameters:
//   - slice: the list of elements to filter.
//
// Returns:
//   - []*T: the list of elements that do not satisfy the predicate. Nil if the
//     list is empty or all elements are nil.
func RejectNils[T any](slice []*T) []*T {
	if len(slice) == 0 {
		return nil
	}

	var count int

	for i := range slice {
		if slice[i] != nil {
			count++
		}
	}

	if count == 0 {
		return nil
	}

	new_slice := make([]*T, 0, count)

	for i := range slice {
		if slice[i] != nil {
			new_slice = append(new_slice, slice[i])
		}
	}

	return new_slice
}
