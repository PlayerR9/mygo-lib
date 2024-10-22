package runes

import "errors"

var (
	// ErrBadEncoding occurs when an input (slice of bytes or string) is not valid utf-8.
	// This error can be checked using the == operator.
	//
	// Format:
	// 	"invalid utf-8"
	ErrBadEncoding error
)

func init() {
	ErrBadEncoding = errors.New("invalid utf-8")
}
