package errors

import (
	"strconv"
	"testing"
)

// TestNew tests New.
func TestNew(t *testing.T) {
	const (
		Want string = "foo"
	)

	err := New(Want)
	if err == nil {
		t.Errorf("want %s, got nothing", strconv.Quote(Want))
	} else {
		msg := err.Error()
		if msg != Want {
			t.Errorf("want %s, got %s", strconv.Quote(Want), strconv.Quote(msg))
		}
	}
}

// TestRejectNilErrors tests RejectNilErrors.
func TestRejectNilErrors(t *testing.T) {
	errs := []error{nil, nil, nil}

	res := RejectNilErrors(errs)
	if res != nil {
		t.Errorf("want nil, got %v", res)
	}
}
