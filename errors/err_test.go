package errors

import (
	"testing"

	test "github.com/PlayerR9/go-verify/test"
)

// TestNewErr tests the NewErr function.
func TestNewErr(t *testing.T) {
	type args struct {
		code     ErrorCode
		msg      string
		expected string
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			err := NewErr(args.code, args.msg)

			err = test.CheckErr(args.expected, err)
			if err != nil {
				t.Error(err)
			}
		}
	})

	_ = tests.AddTest("with message", args{
		code:     BadParameter,
		msg:      "test is invalid",
		expected: "(BadParameter) test is invalid",
	})

	_ = tests.AddTest("without message", args{
		code:     BadParameter,
		msg:      "",
		expected: "(BadParameter) something went wrong",
	})

	_ = tests.AddTest("with invalid code", args{
		code:     -1,
		msg:      "",
		expected: "(ErrorCode(-1)) something went wrong",
	})

	_ = tests.Run(t)
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
