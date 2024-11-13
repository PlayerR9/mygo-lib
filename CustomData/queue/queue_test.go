package queue

import (
	"fmt"
	"slices"
	"testing"
)

func testEnqueue[T comparable](queue Queue[T], want []T) error {
	for _, n := range want {
		err := queue.Enqueue(n)
		if err != nil {
			return fmt.Errorf("queue.Enqueue(%v) = %w", n, err)
		}
	}

	var got []T

	for !queue.IsEmpty() {
		n, err := queue.Dequeue()
		if err != nil {
			return fmt.Errorf("queue.Dequeue() = %w", err)
		}

		got = append(got, n)
	}

	ok := slices.Equal(want, got)
	if !ok {
		return fmt.Errorf("want %v, got %v", want, got)
	}

	return nil
}

func TestQueueEnqueue(t *testing.T) {
	const (
		MAX uint = 10
	)

	want := make([]uint, 0, MAX)
	for i := uint(0); i < MAX; i++ {
		want = append(want, i)
	}

	var queue Queue[uint]

	// 1. Test ArrayQueue
	queue = new(ArrayQueue[uint])

	err := testEnqueue(queue, want)
	if err != nil {
		t.Errorf("ArrayQueue[int] = %s", err.Error())
	}

	// 2. Test LinkedQueue
	queue = new(LinkedQueue[uint])

	err = testEnqueue(queue, want)
	if err != nil {
		t.Errorf("LinkedQueue[int] = %s", err.Error())
	}
}
