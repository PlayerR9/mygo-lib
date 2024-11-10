package list

import "github.com/PlayerR9/mygo-lib/common"

// List is an interface defining a list.
type List[T any] interface {
	// Enlist adds an element at the end of the list.
	//
	// Parameters:
	//   - elem: The element to add.
	//
	// Returns:
	//   - error: An error if the element could not be added to the list.
	//
	// Errors:
	// 	- common.ErrNilReceiver: If the receiver is nil.
	// 	- ErrFullList: If the list has a capacity and that capacity has been reached.
	//    - any other error that may be returned.
	Enlist(elem T) error

	// Prepend adds an element at the start of the list.
	//
	// Parameters:
	//   - elem: The element to add.
	//
	// Returns:
	//   - error: An error if the element could not be added to the list.
	//
	// Errors:
	// 	- common.ErrNilReceiver: If the receiver is nil.
	// 	- ErrFullList: If the list has a capacity and that capacity has been reached.
	//    - any other error that may be returned.
	Prepend(elem T) error

	// Delist removes the first element from the list.
	//
	// Returns:
	//   - T: The element that was removed.
	//   - error: An error if the element could not be removed from the list.
	//
	// Errors:
	//   - ErrEmptyList: If the list is empty.
	//   - common.ErrNilReceiver: If the receiver is nil.
	//   - any other error that may be returned.
	Delist() (T, error)

	// Deback removes the last element from the list.
	//
	// Returns:
	//   - T: The element that was removed.
	//   - error: An error if the element could not be removed from the list.
	//
	// Errors:
	//   - ErrEmptyList: If the list is empty.
	//   - common.ErrNilReceiver: If the receiver is nil.
	//   - any other error that may be returned.
	Deback() (T, error)

	// Front returns the element at the start of the list without removing it.
	//
	// Returns:
	//   - T: The element at the start of the list.
	//   - error: An error if the element could not be peeked from the list.
	//
	// Errors:
	//   - common.ErrNilReceiver: If the receiver is nil.
	//   - common.ErrEmptyList: If the list is empty.
	//   - any other error that may be returned.
	Front() (T, error)

	// Back returns the element at the end of the list without removing it.
	//
	// Returns:
	//   - T: The element at the end of the list.
	//   - error: An error if the element could not be peeked from the list.
	//
	// Errors:
	//   - common.ErrNilReceiver: If the receiver is nil.
	//   - common.ErrEmptyList: If the list is empty.
	//   - any other error that may be returned.
	Back() (T, error)

	// Size returns the number of elements in the list.
	//
	// Returns:
	//   - uint: The number of elements in the list.
	//
	// If the receiver is nil, then 0 is returned.
	Size() uint

	// IsEmpty checks whether the list is empty.
	//
	// Returns:
	//   - bool: True if the list is empty, false otherwise.
	//
	// If the receiver is nil, then true is returned.
	IsEmpty() bool
}

// Enlist adds multiple elements to the list in the order they are passed. If the list implements
// the `EnlistMany` method, then that method is used instead.
//
// Parameters:
//   - list: The list to which the elements are added.
//   - elems: Variadic parameters representing the elements to be added.
//
// Returns:
//   - uint: The number of elements successfully enlistd onto the list.
//   - error: An error if the list is nil or if there is an issue enlisting one of the elements.
//
// Errors:
//   - common.ErrNilParam: If the list is nil.
//   - ErrFullList: If not all elements could be enlistd onto the list.
//   - any error returned by the `Enlist()` method of the list.
func Enlist[T any](list List[T], elems ...T) (uint, error) {
	lenElems := uint(len(elems))
	if lenElems == 0 {
		return 0, nil
	} else if list == nil {
		return 0, common.NewErrNilParam("list")
	}

	l, ok := list.(interface{ EnlistMany(elems []T) (uint, error) })
	if ok {
		n, err := l.EnlistMany(elems)
		return n, err
	}

	for i, elem := range elems {
		err := list.Enlist(elem)
		if err != nil {
			return uint(i), err
		}
	}

	return lenElems, nil
}

// Prepend adds multiple elements before the start of the list in the order they are passed.
// If the list implements the `PrependMany` method, then that method is used instead.
//
// Parameters:
//   - list: The list to which the elements are added.
//   - elems: Variadic parameters representing the elements to be added.
//
// Returns:
//   - uint: The number of elements successfully prepended onto the list.
//   - error: An error if the list is nil or if there is an issue prepending one of the elements.
//
// Errors:
//   - common.ErrNilParam: If the list is nil.
//   - ErrFullList: If not all elements could be prepended onto the list.
//   - any error returned by the `Enlist()` method of the list.
func Prepend[T any](list List[T], elems ...T) (uint, error) {
	lenElems := uint(len(elems))
	if lenElems == 0 {
		return 0, nil
	} else if list == nil {
		return 0, common.NewErrNilParam("list")
	}

	l, ok := list.(interface{ PrependMany(elems []T) (uint, error) })
	if ok {
		n, err := l.PrependMany(elems)
		return n, err
	}

	for i := lenElems - 1; i >= 0; i-- {
		err := list.Prepend(elems[i])
		if err != nil {
			return lenElems - i, err
		}
	}

	return lenElems, nil
}

// Free frees the list. If the list implements `Type` interface, then its `Free()`
// method is called. If not, then the list is cleared by delisting all elements from the list.
//
// Parameters:
//   - list: The list to free.
func Free[T any](list List[T]) {
	if list == nil {
		return
	}

	if l, ok := list.(common.Freeable); ok {
		l.Free()
		return
	}

	for {
		_, err := list.Delist()
		if err != nil {
			break
		}
	}
}

// Reset resets the list for reuse. If the list implements `Resetter` interface,
// then its `Reset()` method is called. If not, then the list is cleared by delisting all
// elements from the list.
//
// Parameters:
//   - list: The list to reset.
func Reset[T any](list List[T]) {
	if list == nil || common.Reset(list) {
		return
	}

	for {
		_, err := list.Delist()
		if err != nil {
			break
		}
	}
}
