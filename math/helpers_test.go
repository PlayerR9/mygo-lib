package math

import (
	"slices"
	"testing"
)

// TestUintPow tests the UintPow function.
func TestUintPow(t *testing.T) {
	type args struct {
		base uint
		exp  uint
	}

	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			"10^4 = 10000",
			args{
				base: 10,
				exp:  4,
			},
			10000,
		},
		{
			"9^3 = 729",
			args{
				base: 9,
				exp:  3,
			},
			729,
		},
	}

	for _, tt := range tests {
		fn := func(t *testing.T) {
			res, err := UintPow(tt.args.base, tt.args.exp)
			if err != nil {
				t.Errorf("want no error, got %v", err)
			} else if res != tt.want {
				t.Errorf("want %d, got %d", tt.want, res)
			}
		}

		_ = t.Run(tt.name, fn)
	}
}

// TestUintPowSlice tests the UintPowSlice function.
func TestUintPowSlice(t *testing.T) {
	type args struct {
		base    uint
		max_exp uint
	}

	tests := []struct {
		name string
		args args
		want []uint
	}{
		{
			"10^4 = [1, 10, 100, 1000, 10000]",
			args{
				base:    10,
				max_exp: 4,
			},
			[]uint{1, 10, 100, 1000, 10000},
		},
		{
			"2^8 = [1, 2, 4, 8, 16, 32, 64, 128, 256]",
			args{
				base:    2,
				max_exp: 8,
			},
			[]uint{1, 2, 4, 8, 16, 32, 64, 128, 256},
		},
	}

	for _, tt := range tests {
		fn := func(t *testing.T) {
			got, err := UintPowSlice(tt.args.base, tt.args.max_exp)
			if err != nil {
				t.Errorf("want no error, got %v", err)
			} else if !slices.Equal(got, tt.want) {
				t.Errorf("want %v, got %v", tt.want, got)
			}
		}

		_ = t.Run(tt.name, fn)
	}
}
