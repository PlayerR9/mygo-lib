package internal

import (
	"slices"
	"testing"
)

// TestSplit tests the Split function.
func TestSplit(t *testing.T) {
	result := Split([]rune("a,b,,c,d"), ',')

	expected := [][]rune{
		[]rune("a"),
		[]rune("b"),
		[]rune("c"),
		[]rune("d"),
	}

	if len(result) != len(expected) {
		t.Errorf("expected %d, got %d", len(expected), len(result))
	} else {
		for i, r := range result {
			if !slices.Equal(r, expected[i]) {
				t.Errorf("expected %v, got %v", expected[i], r)
			}
		}
	}
}
