package pointers

import (
	"testing"

	vc "github.com/PlayerR9/go-verify/common"
	"github.com/PlayerR9/go-verify/test"
)

// TestSet tests the Set function.
func TestSet(t *testing.T) {
	x := 15

	Set(&x, 5)

	if x == 5 {
		return
	}

	vc.FAIL.WrongInt(t, 5, x)
}

// TestGet tests the Get function.
func TestGet(t *testing.T) {
	type args struct {
		x          int
		px         *int
		is_pointer bool
		want       int
	}

	tests := test.NewTests[args](func(args args) test.TestingFunc {
		return func(t *testing.T) {
			var v int

			if args.is_pointer {
				v = Get(args.px)
			} else {
				v = Get(&args.x)
			}

			if v == args.want {
				return
			}

			vc.FAIL.WrongInt(t, args.want, v)
		}
	})

	_ = tests.AddTest("nil pointer", args{
		px:         nil,
		is_pointer: true,
		want:       0,
	})

	_ = tests.AddTest("non-nil pointer", args{
		x:          15,
		is_pointer: false,
		want:       15,
	})

	_ = tests.Run(t)
}
