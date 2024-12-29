package internal

import "testing"

// TestFirstIndexOf tests FirstIndexOf.
func TestFirstIndexOf(t *testing.T) {
	elems := make([]int, 0, 1000)
	for i := 0; i < 1000; i++ {
		elems = append(elems, i)
	}

	idx, ok := FirstIndexOf(elems, func(e int) bool {
		return e > 100
	})

	if !ok {
		t.Errorf("expected index to be present")
	} else if idx != 101 {
		t.Errorf("expected index to be 101, got %d", idx)
	}
}
