package errors

import (
	"strconv"
	"testing"

	"github.com/PlayerR9/mygo-lib/common"
)

// TestThrowTry tests Throw and Try.
func TestThrowTry(t *testing.T) {
	panic_fn := func() {
		Throw(common.DefaultError)
	}

	err := Try(panic_fn)
	if err == common.DefaultError {
		return
	}

	var want string

	if common.DefaultError == nil {
		want = "something"
	} else {
		want = common.DefaultError.Error()
	}

	var got string

	if err == nil {
		got = "nothing"
	} else {
		msg := err.Error()
		got = strconv.Quote(msg)
	}

	t.Errorf("want %s, got %s", want, got)
}
