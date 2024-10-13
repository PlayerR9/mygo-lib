package errors

import (
	"testing"

	test "github.com/PlayerR9/go-verify/test"
	faults "github.com/PlayerR9/mygo-lib/PlayerR9/go-fault"
)

// TestInvalidErrorCode tests what happens when an invalid error code is passed.
func TestInvalidErrorCode(t *testing.T) {
	const (
		Expected string = "(ErrorCode(-1)) something went wrong"
	)

	fault := faults.New[ErrorCode](-1, "something went wrong")
	if fault == nil {
		t.Fatal("want Fault, got nil")
	}

	msg := fault.Error()
	if msg != Expected {
		t.Fatalf("want %q, got %q", Expected, msg)
	}
}

// TestNewBadParameter tests the NewBadParameter function.
func TestNewBadParameter(t *testing.T) {
	type args struct {
		param_name string
		must       string
		expected   string
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			fault := NewBadParameter(args.param_name, args.must)

			err := test.CheckErr(args.expected, fault)
			if err != nil {
				t.Error(err)
			}
		}
	})

	_ = tests.AddTest("no parameter and must", args{
		param_name: "",
		must:       "",
		expected:   "(BadParameter) Parameter is invalid",
	})

	_ = tests.AddTest("no parameter but with must", args{
		param_name: "",
		must:       "be positive",
		expected:   "(BadParameter) Parameter must be positive",
	})

	_ = tests.AddTest("with parameter but no must", args{
		param_name: "x",
		must:       "",
		expected:   "(BadParameter) Parameter \"x\" is invalid",
	})

	_ = tests.AddTest("with parameter and must", args{
		param_name: "x",
		must:       "be positive",
		expected:   "(BadParameter) Parameter \"x\" must be positive",
	})

	_ = tests.Run(t)
}
