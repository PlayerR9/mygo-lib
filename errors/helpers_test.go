package errors

import (
	"testing"
)

// TestRejectNilErrors tests RejectNilErrors.
func TestRejectNilErrors(t *testing.T) {
	errs := []error{nil, nil, nil}

	res := RejectNilErrors(errs)
	if res != nil {
		t.Errorf("want nil, got %v", res)
	}
}
