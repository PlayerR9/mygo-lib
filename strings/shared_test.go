package strings

import (
	"slices"
	"testing"

	vc "github.com/PlayerR9/go-verify/common"
	"github.com/PlayerR9/go-verify/test"
)

// TestIndicesOf tests the IndicesOf function.
func TestIndicesOf(t *testing.T) {
	type args struct {
		slice    []string
		sep      string
		expected []uint
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			indices := IndicesOf(args.slice, args.sep)

			ok := slices.Equal(indices, args.expected)
			if !ok {
				vc.FAIL.WrongAny(t, args.expected, indices)
			}
		}
	})

	_ = tests.AddTest("success", args{
		slice:    []string{"a", "b", "c", "a"},
		sep:      "a",
		expected: []uint{0, 3},
	})

	_ = tests.AddTest("no match", args{
		slice:    []string{"a", "b", "c", "d"},
		sep:      "f",
		expected: nil,
	})

	_ = tests.Run(t)
}
