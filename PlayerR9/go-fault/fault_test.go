package faults

import (
	"testing"

	"github.com/PlayerR9/go-verify/test"
)

type MockIs struct {
	Fault
}

func (mi MockIs) Embeds() Fault {
	return mi.Fault
}

// TestIs tests the Is function.
func TestIs(t *testing.T) {
	type args struct {
		fault, target Fault
		expected      bool
	}

	fault1 := New(MockFaultCodeA, "foo")
	fault2 := New(MockFaultCodeB, "bar")
	fault3 := &MockIs{
		Fault: New(MockFaultCodeA, "foo"),
	}
	fault4 := &MockIs{
		Fault: nil,
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			ok := Is(args.fault, args.target)
			if ok != args.expected {
				t.Errorf("want %t, got %t", args.expected, ok)
			}
		}
	})

	_ = tests.AddTest("embedded fail case", args{
		fault:    fault3,
		target:   fault4,
		expected: false,
	})

	_ = tests.AddTest("nil parameters", args{
		fault:    nil,
		target:   nil,
		expected: false,
	})

	_ = tests.AddTest("self-identity", args{
		fault:    fault1,
		target:   fault1,
		expected: true,
	})

	_ = tests.AddTest("fail case", args{
		fault:    fault1,
		target:   fault2,
		expected: false,
	})

	_ = tests.AddTest("success case", args{
		fault:    fault1,
		target:   fault3,
		expected: true,
	})

	_ = tests.AddTest("embedded case", args{
		fault:    fault3,
		target:   fault2,
		expected: false,
	})

	_ = tests.AddTest("embedded nil case", args{
		fault:    fault4,
		target:   fault1,
		expected: false,
	})

	_ = tests.Run(t)
}
