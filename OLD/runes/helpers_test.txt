package runes

import (
	"slices"
	"testing"

	"github.com/PlayerR9/go-verify/test"
)

// TestQuoteRunes tests the QuoteRunes function.
func TestQuoteRunes(t *testing.T) {
	type args struct {
		slice    []rune
		expected []string
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			quoted := QuoteRunes(args.slice)

			ok := slices.Equal(quoted, args.expected)
			if !ok {
				t.Errorf("expected %v, got %v", args.expected, quoted)
			}
		}
	})

	_ = tests.AddTest("with empty slice", args{
		slice:    []rune{},
		expected: nil,
	})

	_ = tests.AddTest("with slice", args{
		slice:    []rune{'t', 'e', 's', 't'},
		expected: []string{"'t'", "'e'", "'s'", "'t'"},
	})

	_ = tests.Run(t)
}

// TestRunesToStrings tests the RunesToStrings function.
func TestRunesToStrings(t *testing.T) {
	type args struct {
		slice    []rune
		expected []string
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			strs := RunesToStrings(args.slice)

			ok := slices.Equal(strs, args.expected)
			if !ok {
				t.Errorf("expected %v, got %v", args.expected, strs)
			}
		}
	})

	_ = tests.AddTest("with empty slice", args{
		slice:    []rune{},
		expected: nil,
	})

	_ = tests.AddTest("with slice", args{
		slice:    []rune{'t', 'e', 's', 't'},
		expected: []string{"t", "e", "s", "t"},
	})

	_ = tests.Run(t)
}
