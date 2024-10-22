package runes

import (
	"fmt"
	"testing"

	"github.com/PlayerR9/go-verify/test"
)

// TestErrBadEncoding tests the ErrBadEncoding function.
func TestErrBadEncoding(t *testing.T) {
	type args struct {
		idx      int
		expected string
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			err := NewErrBadEncoding(args.idx)

			err = test.CheckErr(args.expected, err)
			if err != nil {
				t.Error(err)
			}
		}
	})

	_ = tests.AddTest("with positive index", args{
		idx:      0,
		expected: fmt.Sprintf("byte %d is not valid utf-8", 0),
	})

	_ = tests.AddTest("with negative index", args{
		idx:      -1,
		expected: fmt.Sprintf("byte %d is not valid utf-8", -1),
	})

	_ = tests.Run(t)
}

// TestNewErrNotAsExpected tests the NewErrNotAsExpected function.
func TestNewErrNotAsExpected(t *testing.T) {
	type args struct {
		previous  rune
		expecteds []rune
		got       *rune
		expected  string
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			res := NewErrNotAsExpected(args.previous, args.expecteds, args.got)

			err := test.CheckErr(args.expected, res)
			if err != nil {
				t.Error(err)
			}
		}
	})

	_ = tests.AddTest("without got", args{
		previous:  'a',
		expecteds: []rune{'a', 'b'},
		got:       nil,
		expected:  "after 'a': expected either 'a' or 'b', got nothing instead",
	})

	var (
		A rune = 'a'
	)

	_ = tests.AddTest("with got", args{
		previous:  'a',
		expecteds: []rune{'a', 'b'},
		got:       &A,
		expected:  "after 'a': expected either 'a' or 'b', got 'a' instead",
	})

	_ = tests.Run(t)
}
