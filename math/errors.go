package math

import "errors"

var (
	// ErrInvalidBase occurs when the base is zero.
	//
	// Format:
	//   "base must not be zero"
	ErrInvalidBase error
)

func init() {
	ErrInvalidBase = errors.New("base must not be zero")
}
