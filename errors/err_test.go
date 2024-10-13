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

// TestNewErrBadParameter tests the NewErrBadParameter function.
func TestNewErrBadParameter(t *testing.T) {
	type args struct {
		param    string
		msg      string
		expected string
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			err := NewErrBadParameter(args.param, args.msg)

			err = test.CheckErr(args.expected, err)
			if err != nil {
				t.Error(err)
			}
		}
	})

	_ = tests.AddTest("with message", args{
		param:    "test",
		msg:      "not be invalid",
		expected: "(BadParameter) parameter (\"test\") must not be invalid",
	})

	_ = tests.AddTest("without message", args{
		param:    "test",
		msg:      "",
		expected: "(BadParameter) parameter (\"test\") is invalid",
	})

	_ = tests.AddTest("empty param", args{
		param:    "",
		msg:      "",
		expected: "(BadParameter) parameter is invalid",
	})

	_ = tests.AddTest("empty param and message", args{
		param:    "",
		msg:      "be greater than 0",
		expected: "(BadParameter) parameter must be greater than 0",
	})

	_ = tests.Run(t)
}

// TestNewErrNilParameter tests the NewErrNilParameter function.
func TestNewErrNilParameter(t *testing.T) {
	type args struct {
		param    string
		expected string
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			err := NewErrNilParameter(args.param)

			err = test.CheckErr(args.expected, err)
			if err != nil {
				t.Error(err)
			}
		}
	})

	_ = tests.AddTest("with param", args{
		param:    "test",
		expected: "(BadParameter) parameter (\"test\") must not be nil",
	})

	_ = tests.AddTest("without param", args{
		param:    "",
		expected: "(BadParameter) parameter must not be nil",
	})

	_ = tests.Run(t)
}
