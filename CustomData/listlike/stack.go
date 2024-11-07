package listlike

import (
	"github.com/PlayerR9/mygo-lib/common"
)

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

	// IsEmpty checks whether the stack is empty.
	//
	// Returns:
	//   - bool: True if the stack is empty, false otherwise.
	IsEmpty() bool
}

// Reset resets the stack for reuse. Does nothing if the parameter is nil.
//
// Parameters:
//   - stack: The stack to be reset.
//
// Panics:
//   - common.ErrMethodNotImpl: If the stack does not implement the Reset() method.
func Reset[T any](stack Stack[T]) {
	if stack == nil {
		return
	}

	s, ok := stack.(interface{ Reset() })
	if !ok {
		panic(common.NewErrMethodNotImpl("stack.Reset()", stack))
	}

	s.Reset()
}

// Peek calls the `Peek()` method on the given stack if it has one, or pops the top
// element, pushes it back and returns it if not.
//
// Parameters:
//   - stack: The stack to peek.
//
// Returns:
//   - any: The top element of the stack. Nil if the stack is empty.
//   - error: An error if the element could not be peeked.
//
// Errors:
//   - common.ErrNilParam: If the stack is nil.
//   - ErrEmptyStack: If the stack is empty.
//   - any other error returned by the `Peek()` or `Pop()` methods.
func Peek[T any](stack Stack[T]) (T, error) {
	if stack == nil {
		return *new(T), common.NewErrNilParam("stack")
	}

	s, ok := stack.(interface{ Peek() (T, error) })
	if ok {
		return s.Peek()
	}

	top, err := stack.Pop()
	if err != nil {
		return *new(T), err
	}

	err = stack.Push(top)
	if err != nil {
		panic(NewErrBadImplement("stack.Push()", err))
	}

	return top, nil
}

// Push adds multiple elements to the stack. If a stack supports the `PushMany“
// method, it is used to add all elements at once.
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
//   - Any error returned by the Push or PushMany methods of the stack.
//
// Behaviours:
//   - Elements are added in reverse order. This means that the first element in the slice
//     is also the first one got upon calling the `Pop()` method.
func Push[T any](stack Stack[T], elems ...T) (uint, error) {
	lenElems := uint(len(elems))
	if lenElems == 0 {
		return 0, nil
	} else if stack == nil {
		return 0, common.NewErrNilParam("stack")
	}

	s, ok := stack.(interface{ PushMany([]T) (uint, error) })
	if ok {
		return s.PushMany(elems)
	}

	for i := lenElems - 1; i >= 0; i-- {
		err := stack.Push(elems[i])
		if err != nil {
			return lenElems - 1 - i, err
		}
	}

	return lenElems, nil
}

/*
func Slice[T any](stack Stack[T]) []T {
	for {
		elem, err := stack.Pop()
		if err == ErrEmptyStack {
			break
		} else if err != nil {

		}

		if err != nil {
			break
		}
	}
} */
