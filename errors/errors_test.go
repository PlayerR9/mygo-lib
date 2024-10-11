package errors

import (
	"errors"
	"testing"

	"github.com/PlayerR9/go-verify/test"
)

// TestNewErrAfter tests the NewErrAfter function.
func TestNewErrAfter(t *testing.T) {
	type args struct {
		previous string
		inner    error
		expected string
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			res := NewErrAfter(args.previous, args.inner)

			err := test.CheckErr(args.expected, res)
			if err != nil {
				t.Error(err)
			}
		}
	})

	_ = tests.AddTest("with previous", args{
		previous: "foo",
		inner:    errors.New("bar"),
		expected: "after foo: bar",
	})

	_ = tests.AddTest("without previous", args{
		previous: "",
		inner:    errors.New("bar"),
		expected: "after something: bar",
	})

	_ = tests.AddTest("without inner", args{
		previous: "foo",
		inner:    nil,
		expected: "after foo: something went wrong",
	})

	_ = tests.Run(t)
}

// TestNewErrAfterUnwrap tests the Unwrap method of the ErrAfter type.
func TestNewErrAfterUnwrap(t *testing.T) {
	var (
		Inner error = errors.New("inner")
	)

	err := NewErrAfter("foo", Inner)
	if err == nil {
		t.Error("want error, got nil")
	} else {
		inner := errors.Unwrap(err)
		if inner != Inner {
			t.Errorf("want %v, got %v", Inner, inner)
		}
	}
}

// TestNewErrNotAsExpected tests the NewErrNotAsExpected function.
func TestNewErrNotAsExpected(t *testing.T) {
	type args struct {
		expecteds []string
		got       *string
		expected  string
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			res := NewErrNotAsExpected(args.expecteds, args.got)

			err := test.CheckErr(args.expected, res)
			if err != nil {
				t.Error(err)
			}
		}
	})

	_ = tests.AddTest("with expecteds but without got", args{
		expecteds: []string{"foo", "bar"},
		got:       nil,
		expected:  "expected either foo or bar, got nothing instead",
	})

	_ = tests.AddTest("without expecteds and got", args{
		expecteds: nil,
		got:       nil,
		expected:  "expected nothing, got nothing instead",
	})

	var (
		Baz string = "baz"
	)

	_ = tests.AddTest("with expecteds and got", args{
		expecteds: []string{"foo", "bar"},
		got:       &Baz,
		expected:  "expected either foo or bar, got baz instead",
	})

	_ = tests.AddTest("with expecteds but without got", args{
		expecteds: []string{"foo", "bar"},
		got:       nil,
		expected:  "expected either foo or bar, got nothing instead",
	})

	_ = tests.Run(t)
}
