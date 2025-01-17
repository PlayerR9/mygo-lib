package runes

import "errors"

var (
	// ErrInvalidUtf8 occurs when a certain byte sequence is not valid utf-8. This
	// error can be checked with the == operator.
	//
	// Format:
	// 	"invalid utf-8"
	ErrInvalidUtf8 error

	// ErrNoPredicate occurs when a predicate is not provided.
	//
	// Format:
	// 	"no predicate was provided"
	ErrNoPredicate error
)

func init() {
	ErrInvalidUtf8 = errors.New("invalid utf-8")
	ErrNoPredicate = errors.New("no predicate was provided")
}
