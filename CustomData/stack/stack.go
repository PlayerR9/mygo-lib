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
	//   - T: The element that was removed.
	//   - error: An error if the element could not be removed from the stack.
	//
	// Errors:
	//   - ErrEmptyStack: If the stack is empty.
	//   - common.ErrNilReceiver: If the receiver is nil.
	//   - any other error that may be returned.
	Pop() (T, error)

	// Peek returns the element at the top of the stack without removing it.
	//
	// Returns:
	//   - T: The element at the top of the stack.
	//   - error: An error if the element could not be peeked from the stack.
	//
	// Errors:
	//   - common.ErrNilReceiver: If the receiver is nil.
	//   - common.ErrEmptyStack: If the stack is empty.
	//   - any other error that may be returned.
	Peek() (T, error)

	// Size returns the number of elements in the stack.
	//
	// Returns:
	//   - uint: The number of elements in the stack.
	//
	// If the receiver is nil, then 0 is returned.
	Size() uint

	// IsEmpty checks whether the stack is empty.
	//
	// Returns:
	//   - bool: True if the stack is empty, false otherwise.
	//
	// If the receiver is nil, then true is returned.
	IsEmpty() bool
}

// Push adds multiple elements to the stack in reverse order. If the stack implements
// the `PushMany` method, then that method is used instead.
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

	s, ok := stack.(interface{ PushMany(elems []T) (uint, error) })
	if ok {
		n, err := s.PushMany(elems)
		return n, err
	}

	for i := lenElems - 1; i >= 0; i-- {
		err := stack.Push(elems[i])
		if err != nil {
			return lenElems - i, err
		}
	}

	return lenElems, nil
}

// Free frees the stack. If the stack implements `Type` interface, then its `Free()`
// method is called. If not, then the stack is cleared by popping all elements from the stack.
//
// Parameters:
//   - stack: The stack to free.
func Free[T any](stack Stack[T]) {
	if stack == nil {
		return
	}

	if s, ok := stack.(common.Freeable); ok {
		s.Free()
		return
	}

	for {
		_, err := stack.Pop()
		if err != nil {
			break
		}
	}
}

// Reset resets the stack for reuse. If the stack implements `Resetter` interface,
// then its `Reset()` method is called. If not, then the stack is cleared by popping all
// elements from the stack.
//
// Parameters:
//   - stack: The stack to reset.
func Reset[T any](stack Stack[T]) {
	if stack == nil || common.Reset(stack) {
		return
	}

	for {
		_, err := stack.Pop()
		if err != nil {
			break
		}
	}
}
