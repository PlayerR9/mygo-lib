package faults

import "testing"

type MockIs struct {
	Fault
}

func (mi MockIs) Embeds() Fault {
	return mi.Fault
}

func (mi MockIs) IsFault(_ Fault) bool {
	return true
}

func TestIs(t *testing.T) {
	fault1 := New(MockFaultCodeA, "foo")
	fault2 := New(MockFaultCodeB, "bar")
	fault3 := &MockIs{}

	ok := Is(nil, nil)
	if ok {
		t.Errorf("want false, got true")
	}

	ok = Is(fault1, fault1)
	if !ok {
		t.Errorf("want true, got false")
	}

	ok = Is(fault3, fault1)
	if !ok {
		t.Errorf("want true, got false")
	}

	ok = Is(fault1, fault2)
	if ok {
		t.Errorf("want false, got true")
	}
}
