package listlike

import (
	"github.com/PlayerR9/mygo-lib/common"
)

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

	// IsEmpty checks whether the stack is empty.
	//
	// Returns:
	//   - bool: True if the stack is empty, false otherwise.
	IsEmpty() bool
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
func Peek(stack Stack) (any, error) {
	if stack == nil {
		return nil, common.NewErrNilParam("stack")
	}

	s, ok := stack.(interface{ Peek() (any, error) })
	if ok {
		return s.Peek()
	}

	top, err := stack.Pop()
	if err != nil {
		return nil, err
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
func Push(stack Stack, elems ...any) (uint, error) {
	lenElems := uint(len(elems))
	if lenElems == 0 {
		return 0, nil
	} else if stack == nil {
		return 0, common.NewErrNilParam("stack")
	}

	s, ok := stack.(interface{ PushMany([]any) (uint, error) })
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
