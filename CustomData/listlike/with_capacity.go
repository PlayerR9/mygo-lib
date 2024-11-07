package listlike

import "github.com/PlayerR9/mygo-lib/common"

type CapacityStack[T any] struct {
	stack    Stack[T]
	size     uint
	capacity uint
}

// IsEmpty implements Stack.
func (c CapacityStack[T]) IsEmpty() bool {
	return c.size == 0
}

// Pop implements Stack.
func (c *CapacityStack[T]) Pop() (T, error) {
	if c == nil {
		return *new(T), common.ErrNilReceiver
	} else if c.size == 0 {
		return *new(T), ErrEmptyStack
	}

	top, err := c.stack.Pop()
	if err != nil {
		panic(err)
	}

	c.size--

	return top, nil
}

// Push implements Stack.
func (c *CapacityStack[T]) Push(elem T) error {
	if c == nil {
		return common.ErrNilReceiver
	} else if c.size == c.capacity {
		return ErrFullStack
	}

	err := c.stack.Push(elem)
	if err != nil {
		panic(err)
	}
	c.size++

	return nil
}

func NewCapacityStack[T any](stack Stack[T], capacity uint) (Stack[T], error) {
	if stack == nil {
		return nil, common.NewErrNilParam("stack")
	}

	return &CapacityStack[T]{
		stack:    stack,
		size: ,
		capacity: capacity,
	}, nil
}
