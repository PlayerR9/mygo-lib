package strings

import (
	"strconv"
	"strings"
)

// ErrNotAsExpected occurs when a string is not as expected.
type ErrNotAsExpected struct {
	// Kind is the kind of the string that is not as expected.
	Kind string

	// Expected are the string that were expected.
	Expecteds []string

	// Got is the actual string.
	Got *string
}

// Error implements the error interface.
func (e ErrNotAsExpected) Error() string {
	var kind string

	if e.Kind != "" {
		kind = e.Kind + " to be "
	}

	var got string

	if e.Got == nil {
		got = "nothing"
	} else {
		got = strconv.Quote(*e.Got)
	}

	var builder strings.Builder

	builder.WriteString("expected ")
	builder.WriteString(kind)

	elems := make([]string, 0, len(e.Expecteds))

	for _, elem := range e.Expecteds {
		elems = append(elems, strconv.Quote(elem))
	}

	builder.WriteString(EitherOrString(elems))
	builder.WriteString(", got ")
	builder.WriteString(got)

	return builder.String()
}

// NewErrNotAsExpected returns a new ErrNotAsExpected error.
//
// Parameters:
//   - idx: The index of the string that is not as expected.
//   - expected: The string that is not as expected.
//   - got: The actual string.
//
// Returns:
//   - error: The new error. Never returns nil.
//
// Format:
//
//	"expected string at index <idx> to be <expected>, got <got>"
//
// Where:
//   - <idx>: The index of the string that is not as expected.
//   - <expected>: The string that is not as expected. This is quoted with strconv.QuoteRune.
//   - <got>: The actual string. This is quoted with strconv.QuoteRune. If nil, "nothing" is used instead.
func NewErrNotAsExpected(kind string, got *string, expecteds ...string) error {
	return &ErrNotAsExpected{
		Kind:      kind,
		Expecteds: expecteds,
		Got:       got,
	}
}
