package common

import (
	"strconv"
	"testing"
)

// TestThrowTry tests Throw and Try.
func TestThrowTry(t *testing.T) {
	panic_fn := func() {
		Throw(DefaultError)
	}

	err := Try(panic_fn)
	if err == DefaultError {
		return
	}

	var want string

	if DefaultError == nil {
		want = "something"
	} else {
		want = DefaultError.Error()
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
