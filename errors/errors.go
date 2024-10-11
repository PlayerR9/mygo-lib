package errors

import (
	"errors"
)

var (
	// ErrNilReceiver is an error that is returned when a receiver is nil.
	ErrNilReceiver error
)

func init() {
	ErrNilReceiver = errors.New("receiver must not be nil")
}
