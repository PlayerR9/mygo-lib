package errors

import (
	"errors"
	"testing"

	"github.com/PlayerR9/go-verify/test"
)

// TestErrMsgOf tests the ErrMsgOf function.
func TestErrMsgOf(t *testing.T) {
	type args struct {
		err      error
		expected string
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			msg := ErrMsgOf(args.err)
			if msg != args.expected {
				t.Errorf("want %q, got %q", args.expected, msg)
			}
		}
	})

	// 1. Test that, if no error is passed, the default error message is returned.
	_ = tests.AddTest("no error", args{
		err:      nil,
		expected: DefaultErr.Error(),
	})

	// 2. Test that, if an error is passed, the expected error message is returned.
	_ = tests.AddTest("error", args{
		err:      errors.New("foo"),
		expected: "foo",
	})

	// 3. Test that, if an empty error is passed, the empty string is returned.
	_ = tests.AddTest("empty error", args{
		err:      errors.New(""),
		expected: "",
	})

	// 4. Test that, if the error's message is the same as the default error message, it will be
	// undistinguishable from a nil error.
	_ = tests.AddTest("default error", args{
		err:      errors.New("something went wrong"),
		expected: DefaultErr.Error(),
	})

	_ = tests.Run(t)
}
