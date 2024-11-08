package stack

import "github.com/PlayerR9/mygo-lib/common"

// Stack is an interface defining a stack.
type Stack[T any] interface {
	// Push adds an element on the top of the stack.
	//
	// Parameters:
	//   - elem: The element to add.
	//
	// Returns:
	//   - error: An error if the element could not be added to the stack.
	//
	// Errors:
	// 	- common.ErrNilReceiver: If the receiver is nil.
	// 	- ErrFullStack: If the stack has a capacity and that capacity has been reached.
	//    - any other error that may be returned.
	Push(elem T) error

	// Pop removes the top element from the stack.
	//
	// Returns:
	//   - T: The element that was removed. Nil if no element was removed.
	//   - error: An error if the element could not be removed from the stack.
	//
	// Errors:
	//   - ErrEmptyStack: If the stack is empty.
	//   - common.ErrNilReceiver: If the receiver is nil.
	Pop() (T, error)

	// Peek returns the element at the top of the stack without removing it.
	//
	// Returns:
	//   - T: The element at the top of the stack. Nil if the stack is empty.
	//   - error: An error of type ErrEmptyStack if the stack is empty.
	Peek() (T, error)

	// Size returns the number of elements in the stack.
	//
	// Returns:
	//   - uint: The number of elements in the stack.
	Size() uint

	// IsEmpty checks whether the stack is empty.
	//
	// Returns:
	//   - bool: True if the stack is empty, false otherwise.
	IsEmpty() bool
}

// Push adds multiple elements to the stack in reverse order. Otherwise, elements are
// added individually starting from the last element.
//
// Parameters:
//   - stack: The stack to which the elements are added.
//   - elems: Variadic parameters representing the elements to be added.
//
// Returns:
//   - uint: The number of elements successfully pushed onto the stack.
//   - error: An error if the stack is nil or if there is an issue pushing one of the elements.
//
// Errors:
//   - common.ErrNilParam: If the stack is nil.
//   - ErrFullStack: If not all elements could be pushed onto the stack.
//   - any error returned by the `Push()` method of the stack.
func Push[T any](stack Stack[T], elems ...T) (uint, error) {
	lenElems := uint(len(elems))
	if lenElems == 0 {
		return 0, nil
	} else if stack == nil {
		return 0, common.NewErrNilParam("stack")
	}

	for i := lenElems; i >= 0; i-- {
		elem := elems[i]

		err := stack.Push(elem)
		if err != nil {
			return lenElems - i, err
		}
	}

	return lenElems, nil
}
