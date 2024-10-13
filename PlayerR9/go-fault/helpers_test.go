package faults

import "testing"

// TestInnermost tests the Innermost function.
func TestInnermost(t *testing.T) {
	inner := Innermost(nil)
	if inner != nil {
		t.Errorf("want nil, got %T", inner)
	}

	base := New(MockFaultCodeA, "foo")
	if base == nil {
		t.Fatalf("want %T, got nil", base)
	}

	inner = Innermost(&MockIs{
		Fault: base,
	})
	if inner != base {
		t.Errorf("want %T, got %T", base, inner)
	}
}
