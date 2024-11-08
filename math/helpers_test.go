package math

import (
	"slices"
	"testing"

	vc "github.com/PlayerR9/go-verify/common"
	"github.com/PlayerR9/go-verify/test"
)

// TestUintPow tests the UintPow function.
func TestUintPow(t *testing.T) {
	type args struct {
		base uint
		exp  uint
		want uint
	}

	tests := test.NewTests[args](func(args args) test.TestingFunc {
		return func(t *testing.T) {
			res, err := UintPow(args.base, args.exp)
			if err != nil {
				vc.FAIL.WrongError(t, "", err)
				return
			}

			if res != args.want {
				vc.FAIL.WrongInt(t, int(args.want), int(res))
			}
		}
	})

	_ = tests.AddTest("10^4 = 10000", args{
		base: 10,
		exp:  4,
		want: 10000,
	})

	_ = tests.AddTest("9^3 = 729", args{
		base: 9,
		exp:  3,
		want: 729,
	})

	_ = tests.Run(t)
}

// TestUintPowSlice tests the UintPowSlice function.
func TestUintPowSlice(t *testing.T) {
	type args struct {
		base    uint
		max_exp uint
		want    []uint
	}

	tests := test.NewTests[args](func(args args) test.TestingFunc {
		return func(t *testing.T) {
			res, err := UintPowSlice(args.base, args.max_exp)
			if err != nil {
				vc.FAIL.WrongError(t, "", err)
				return
			}

			if !slices.Equal(res, args.want) {
				vc.FAIL.WrongAny(t, args.want, res)
			}
		}
	})

	_ = tests.AddTest("10^4 = [1, 10, 100, 1000, 10000]", args{
		base:    10,
		max_exp: 4,
		want:    []uint{1, 10, 100, 1000, 10000},
	})

	_ = tests.AddTest("2^8 = [1, 2, 4, 8, 16, 32, 64, 128, 256]", args{
		base:    2,
		max_exp: 8,
		want:    []uint{1, 2, 4, 8, 16, 32, 64, 128, 256},
	})

	_ = tests.Run(t)
}
