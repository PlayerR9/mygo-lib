package errors

import (
	"strconv"
	"testing"

	"github.com/PlayerR9/mygo-lib/common"
)

// TestThrowTry tests Throw and Try.
func TestThrowTry(t *testing.T) {
	ErrFoo := common.New("foo")

	panic_fn := func() {
		panic(ErrFoo)
	}

	err := Try(panic_fn)
	if err == ErrFoo {
		return
	}

	want := ErrFoo.Error()
	want = strconv.Quote(want)

	var got string

	if err == nil {
		got = "nothing"
	} else {
		msg := err.Error()
		got = strconv.Quote(msg)
	}

	t.Errorf("want %s, got %s", want, got)
}
