package pointers

import (
	"testing"

	"github.com/PlayerR9/go-verify/test"
)

// TestSet tests the Set function.
func TestSet(t *testing.T) {
	x := 15

	Set(&x, 5)

	if x == 5 {
		return
	}

	err := test.FAIL.Int(5, x)
	t.Fatal(err)
}

// TestGet tests the Get function.
func TestGet(t *testing.T) {
	type args struct {
		x          int
		px         *int
		is_pointer bool
		want       int
	}

	tests := test.NewTestSet[args](func(args args) test.TestingFn {
		return func() error {
			var v int

			if args.is_pointer {
				v = Get(args.px)
			} else {
				v = Get(&args.x)
			}

			if v == args.want {
				return nil
			}

			return test.FAIL.Int(args.want, v)
		}
	})

	_ = tests.Add("nil pointer", args{
		px:         nil,
		is_pointer: true,
		want:       0,
	})

	_ = tests.Add("non-nil pointer", args{
		x:          15,
		is_pointer: false,
		want:       15,
	})

	_ = tests.Run(t)
}
