package runes

import (
	"strconv"

	gers "github.com/PlayerR9/mygo-lib/errors"
)

// ErrBadEncoding occurs when some byte is not valid utf-8.
type ErrBadEncoding struct {
	// Idx is the index of the byte that is not valid utf-8.
	Idx int
}

// Error implements the error interface.
func (e ErrBadEncoding) Error() string {
	return "byte " + strconv.Itoa(e.Idx) + " is not valid utf-8"
}

// NewErrBadEncoding creates a new ErrBadEncoding error.
//
// Parameters:
//   - idx: The index of the byte that is not valid utf-8.
//
// Returns:
//   - error: The new error. Never returns nil.
//
// Format:
//
//	"byte <idx> is not valid utf-8"
//
// Where:
//   - <idx>: The index of the byte that is not valid utf-8.
func NewErrBadEncoding(idx int) error {
	return &ErrBadEncoding{
		Idx: idx,
	}
}

// NewErrNotAsExpected is a convenience function that creates a new errors.ErrAfter error with
// a not as expected error.
//
// Parameters:
//   - previous: The element before the one that caused the error.
//   - expecteds: The expected elements.
//   - got: The actual element.
//
// Returns:
//   - error: The new error. Never returns nil.
func NewErrNotAsExpected(previous rune, expecteds []rune, got *rune) error {
	elems := QuoteRunes(expecteds)

	if got == nil {
		return gers.NewErrAfter(
			strconv.QuoteRune(previous),
			gers.NewErrNotAsExpected(elems, nil),
		)
	} else {
		str := strconv.QuoteRune(*got)

		return gers.NewErrAfter(
			strconv.QuoteRune(previous),
			gers.NewErrNotAsExpected(elems, &str),
		)
	}
}
