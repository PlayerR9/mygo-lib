package runes

import (
	"errors"
	"strconv"
	"strings"
)

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

// ErrNotAsExpected occurs when a rune is not as expected.
type ErrNotAsExpected struct {
	// Idx is the index of the rune that is not as expected.
	Idx int

	// Expected is the rune that is not as expected.
	Expected rune

	// Got is the actual rune.
	Got *rune
}

// Error implements the error interface.
func (e ErrNotAsExpected) Error() string {
	var got string

	if e.Got == nil {
		got = "nothing"
	} else {
		got = strconv.QuoteRune(*e.Got)
	}

	var builder strings.Builder

	builder.WriteString("expected rune at index ")
	builder.WriteString(strconv.Itoa(e.Idx))
	builder.WriteString(" to be ")
	builder.WriteString(strconv.QuoteRune(e.Expected))
	builder.WriteString(", got ")
	builder.WriteString(got)

	return builder.String()
}

// NewErrNotAsExpected returns a new ErrNotAsExpected error.
//
// Parameters:
//   - idx: The index of the rune that is not as expected.
//   - expected: The rune that is not as expected.
//   - got: The actual rune.
//
// Returns:
//   - error: The new error. Never returns nil.
//
// Format:
//
//	"expected rune at index <idx> to be <expected>, got <got>"
//
// Where:
//   - <idx>: The index of the rune that is not as expected.
//   - <expected>: The rune that is not as expected. This is quoted with strconv.QuoteRune.
//   - <got>: The actual rune. This is quoted with strconv.QuoteRune. If nil, "nothing" is used instead.
func NewErrNotAsExpected(idx int, expected rune, got *rune) error {
	return &ErrNotAsExpected{
		Idx:      idx,
		Expected: expected,
		Got:      got,
	}
}
