package faults

import (
	"testing"

	"github.com/PlayerR9/go-verify/test"
)

type MockFaultCode int

const (
	MockFaultCodeA MockFaultCode = iota
	MockFaultCodeB
)

func (m MockFaultCode) String() string {
	return [...]string{
		"MockFaultCodeA",
		"MockFaultCodeB",
	}[m]
}

// TestNew tests the New function.
func TestNew(t *testing.T) {
	fault := New(MockFaultCodeA, "foo")
	if fault == nil {
		t.Fatalf("want %T, got nil", &baseFault[MockFaultCode]{})
	}

	msg := fault.Error()
	if msg != "(MockFaultCodeA) foo" {
		t.Fatalf("want %q, got %q", "(MockFaultCodeA) foo", msg)
	}
}

// TestNewf tests the Newf function.
func TestNewf(t *testing.T) {
	type args struct {
		code    MockFaultCode
		format  string
		args    []any
		message string
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			fault := Newf[MockFaultCode](args.code, args.format, args.args...)
			if fault == nil {
				t.Fatalf("want %T, got nil", &baseFault[MockFaultCode]{})
			}

			msg := fault.Error()
			if msg != args.message {
				t.Fatalf("want %q, got %q", args.message, msg)
			}
		}
	})

	_ = tests.AddTest("without args", args{
		code:    MockFaultCodeA,
		format:  "foo",
		args:    nil,
		message: "(MockFaultCodeA) foo",
	})

	_ = tests.AddTest("with args", args{
		code:    MockFaultCodeA,
		format:  "foo %s",
		args:    []any{"bar"},
		message: "(MockFaultCodeA) foo bar",
	})

	_ = tests.Run(t)
}
