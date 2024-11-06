package listlike

import "github.com/PlayerR9/mygo-lib/common"

// Stack is an interface defining a stack.
type Stack interface {
	// Push adds an element on the top of the stack.
	//
	// Parameters:
	//   - elem: The element to add.
	//
	// Returns:
	//   - error: An error of type common.ErrNilReceiver if the receiver is nil.
	Push(elem any) error

	// Pop removes the top element from the stack.
	//
	// Returns:
	//   - any: The element that was removed. Nil if no element was removed.
	//   - error: An error if the element could not be removed from the stack.
	//
	// Errors:
	//   - ErrEmptyStack: If the stack is empty.
	//   - common.ErrNilReceiver: If the receiver is nil.
	Pop() (any, error)

	// Peek returns the element at the top of the stack without removing it.
	//
	// Returns:
	//   - any: The element at the top of the stack. Nil if the stack is empty.
	//   - error: An error of type ErrEmptyStack if the stack is empty.
	Peek() (any, error)

	// IsEmpty checks whether the stack is empty.
	//
	// Returns:
	//   - bool: True if the stack is empty, false otherwise.
	IsEmpty() bool
}

// PushMany pushes multiple elements onto the stack. If the stack does not have
// a PushMany method, the elements will be pushed one by one.
//
// Parameters:
//   - stack: The stack to push onto.
//   - elems: The elements to push.
//
// Returns:
//   - uint: The number of elements that were pushed.
//   - error: An error if the stack was nil or if an error occurred while pushing
//     the elements.
//
// If the stack does not have a PushMany method and an error occurs while pushing
// the elements, the returned uint is the number of elements that were pushed
// before the error occurred.
func PushMany(stack Stack, elems []any) (uint, error) {
	lenElems := uint(len(elems))
	if len(elems) == 0 {
		return 0, nil
	} else if stack == nil {
		return 0, common.NewErrNilParam("stack")
	}

	s, ok := stack.(interface{ PushMany([]any) (uint, error) })
	if ok {
		return s.PushMany(elems)
	}

	for i := uint(0); i < lenElems; i++ {
		err := stack.Push(elems[i])
		if err != nil {
			return i, err
		}
	}

	return lenElems, nil
}
