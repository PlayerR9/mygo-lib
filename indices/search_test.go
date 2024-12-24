package indices

import "testing"

// TestFirstIndexOf tests FirstIndexOf.
func TestFirstIndexOf(t *testing.T) {
	elems := make([]int, 0, 1000)
	for i := 0; i < 1000; i++ {
		elems = append(elems, i)
	}

	idx_opt := FirstIndexOf(elems, func(e int) bool {
		return e > 100
	})

	ok := idx_opt.IsPresent()
	if !ok {
		t.Errorf("expected index to be present")
	}

	idx := MustGet(idx_opt)

	if idx != 101 {
		t.Errorf("expected index to be 101, got %d", idx)
	}
}
