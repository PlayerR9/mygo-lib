package stack

import (
	"fmt"
	"slices"
	"testing"
)

func testPush[T comparable](stack Stack[T], want []T) error {
	for _, n := range want {
		err := stack.Push(n)
		if err != nil {
			return fmt.Errorf("stack.Push(%v) = %w", n, err)
		}
	}

	var got []T

	for !stack.IsEmpty() {
		n, err := stack.Pop()
		if err != nil {
			return fmt.Errorf("stack.Pop() = %w", err)
		}

		got = append(got, n)
	}

	slices.Reverse(got)

	ok := slices.Equal(want, got)
	if !ok {
		return fmt.Errorf("want %v, got %v", want, got)
	}

	return nil
}

func TestPush(t *testing.T) {
	const (
		MAX uint = 10
	)

	want := make([]uint, 0, MAX)
	for i := uint(0); i < MAX; i++ {
		want = append(want, i)
	}

	var stack Stack[uint]

	// 1. Test ArrayStack
	stack = new(ArrayStack[uint])

	err := testPush(stack, want)
	if err != nil {
		t.Errorf("ArrayStack[int] = %s", err.Error())
	}

	// 2. Test LinkedStack
	stack = new(LinkedStack[uint])

	err = testPush(stack, want)
	if err != nil {
		t.Errorf("LinkedStack[int] = %s", err.Error())
	}
}
