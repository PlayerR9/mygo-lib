package list

import (
	"fmt"
	"slices"
	"testing"
)

func testEnlisting[T comparable](list List[T], want []T) error {
	for _, n := range want {
		err := list.Enlist(n)
		if err != nil {
			return fmt.Errorf("list.Enlist(%v) = %w", n, err)
		}
	}

	var got []T

	for !list.IsEmpty() {
		n, err := list.Delist()
		if err != nil {
			return fmt.Errorf("list.Delist() = %w", err)
		}

		got = append(got, n)
	}

	ok := slices.Equal(want, got)
	if !ok {
		return fmt.Errorf("want %v, got %v", want, got)
	}

	return nil
}

func TestListEnlist(t *testing.T) {
	const (
		MAX uint = 10
	)

	want := make([]uint, 0, MAX)
	for i := uint(0); i < MAX; i++ {
		want = append(want, i)
	}

	var list List[uint]

	// 1. Test ArrayList
	list = new(ArrayList[uint])

	err := testEnlisting(list, want)
	if err != nil {
		t.Errorf("ArrayList[int] = %s", err.Error())
	}

	// 2. Test LinkedList
	list = new(LinkedList[uint])

	err = testEnlisting(list, want)
	if err != nil {
		t.Errorf("LinkedList[int] = %s", err.Error())
	}
}
