package common

import "testing"

// TestGet tests the Get function.
func TestGet(t *testing.T) {
	v := Get[int](nil)
	if v != 0 {
		t.Errorf("want 0, got %d", v)
	}

	x := 15

	v = Get(&x)
	if v != 15 {
		t.Errorf("want 15, got %d", v)
	}
}

// TestSet tests the Set function.
func TestSet(t *testing.T) {
	x := 15

	Set(nil, 5)

	Set(&x, 5)
	if x != 5 {
		t.Errorf("want 5, got %d", x)
	}
}
