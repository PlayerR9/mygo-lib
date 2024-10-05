package slices

import (
	"slices"
	"testing"
)

// TestApplyFilter tests the ApplyFilter function.
func TestApplyFilter(t *testing.T) {
	const (
		MAX int = 100
	)

	expected_evens := make([]int, 0, MAX/2)
	for i := 0; i < MAX; i += 2 {
		expected_evens = append(expected_evens, i)
	}

	nums := make([]int, 0, 100)
	for i := 0; i < MAX; i++ {
		nums = append(nums, i)
	}

	evens := ApplyFilter(nums, func(x int) bool {
		return x%2 == 0
	})

	ok := slices.Equal(evens, expected_evens)
	if !ok {
		t.Errorf("expected %v, got %v", expected_evens, evens)
	}
}

// TestApplyReject tests the ApplyReject function.
func TestApplyReject(t *testing.T) {
	const (
		MAX int = 100
	)

	expected_odds := make([]int, 0, MAX/2)
	for i := 1; i < MAX; i += 2 {
		expected_odds = append(expected_odds, i)
	}

	nums := make([]int, 0, 100)
	for i := 0; i < MAX; i++ {
		nums = append(nums, i)
	}

	odds := ApplyReject(nums, func(x int) bool {
		return x%2 == 0
	})

	ok := slices.Equal(odds, expected_odds)
	if !ok {
		t.Errorf("expected %v, got %v", expected_odds, odds)
	}
}
