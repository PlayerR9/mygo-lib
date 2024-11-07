package listlike

import (
	"slices"
	"testing"
)

func TestPushNoCap(t *testing.T) {
	type args struct {
		Elems []int
	}

	tests := []args{
		{
			Elems: []int{1, 2, 3, 4, 5},
		},
		{
			Elems: nil,
		},
	}

	for _, test := range tests {
		stack := NewArrayStack(NoCapacity, test.Elems...)

		var popped []int

		for !stack.IsEmpty() {
			top, err := stack.Pop()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			popped = append(popped, top)
		}

		if len(popped) != len(test.Elems) {
			t.Fatalf("unexpected length: %v", len(popped))
		}

		ok := slices.Equal(popped, test.Elems)
		if !ok {
			t.Fatalf("unexpected elements: %v", popped)
		}
	}
}

func TestPushWithCap(t *testing.T) {
	elems := []int{1, 2, 3, 4, 5}

	stack := NewArrayStack(5, elems...)

	var count int

	for !stack.IsEmpty() {
		top, err := stack.Pop()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if top != elems[count] {
			t.Fatalf("unexpected element: %v", top)
		}

		count++
	}

	if count != len(elems) {
		t.Fatalf("unexpected count: %v", count)
	}

	Reset(stack)

	elems = []int{1, 2, 3, 4, 5, 6}

	for _, elem := range elems[:len(elems)-1] {
		err := stack.Push(elem)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}

	err := stack.Push(elems[len(elems)-1])
	if err == nil {
		t.Fatal("expected error, got nil")
	} else if err != ErrFullStack {
		t.Fatalf("unexpected error: %v", err)
	}
}

/*
func TestPop(t *testing.T) {

// Pop implements the Stack interface.
func (s *ArrayStack) Pop() (any, error) {
	if s == nil {
		return nil, common.ErrNilReceiver
	} else if len(s.slice) == 0 {
		return nil, ErrEmptyStack
	}

	elem := s.slice[len(s.slice)-1]
	s.slice = s.slice[:len(s.slice)-1]

	return elem, nil
}
}

func TestNewArrayStack(t *testing.T) {

// NewArrayStack creates a new stack from a slice.
//
// Parameters:
//   - cap: The maximum capacity of the stack. If it is equal to NoCapacity, the stack
//     will have no capacity.
//   - elems: The elements to add to the stack.
//
// Returns:
//   - Stack: The new stack. Never returns nil.
//
// If the length of elems is larger than the given capacity, the capacity will be set
// to NoCapacity and the stack will have no capacity.
func NewArrayStack(cap uint8, elems ...any) Stack {
	stack := &ArrayStack{
		slice:    nil,
		capacity: cap,
	}

	_, err := stack.PushMany(elems)
	if err != nil {
		return nil
	}

	return stack
}

}

func TestPushMany(t *testing.T) {

// PushMany adds multiple elements to the stack. If it has a capacity and the total
// length of the stack's underlying slice and the provided slice is larger than the
// capacity, the elements are truncated to fit the capacity.
//
// Parameters:
//   - elems: The elements to add.
//
// Returns:
//   - uint: The number of elements successfully pushed onto the stack.
//   - error: An error if the receiver is nil.
//
// Errors:
//   - common.ErrNilReceiver: If the receiver is nil.
func (s *ArrayStack) PushMany(elems []any) (uint, error) {
	lenElems := uint(len(elems))
	if lenElems == 0 {
		return 0, nil
	} else if s == nil {
		return 0, common.ErrNilReceiver
	}

	originalLen := uint(len(s.slice))
	totalLen := originalLen + lenElems

	var err error

	if s.capacity != NoCapacity && totalLen > uint(s.capacity) {
		lenElems = uint(s.capacity) - originalLen
		err = ErrFullStack
	}

	s.slice = append(s.slice, elems[:lenElems]...)
	slices.Reverse(s.slice[originalLen:])

	return lenElems, err
}

}

func TestReset(t *testing.T) {

// Reset resets the stack for reuse. Does nothing if the receiver is nil.
func (s *ArrayStack) Reset() {
	if s == nil || len(s.slice) == 0 {
		return
	}

	clear(s.slice)

	if s.capacity == NoCapacity {
		s.slice = nil
	} else {
		s.slice = make([]any, 0, s.capacity)
	}
}

}

func TestPeek(t *testing.T) {

// Peek returns the element at the top of the stack without removing it.
//
// Returns:
//   - any: The element at the top of the stack. Nil if the stack is empty.
//   - error: An error of type ErrEmptyStack if the stack is empty.
func (s ArrayStack) Peek() (any, error) {
	if len(s.slice) == 0 {
		return nil, ErrEmptyStack
	}

	return s.slice[len(s.slice)-1], nil
}

}
*/
