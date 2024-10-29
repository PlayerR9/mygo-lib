package strings

import (
	"strconv"
	"strings"

	"github.com/PlayerR9/mygo-lib/common"
)

// ErrNotAsExpected occurs when a string is not as expected.
type ErrNotAsExpected struct {
	// Quote if true, the strings will be quoted before being printed.
	Quote bool

	// Kind is the kind of the string that is not as expected.
	Kind string

	// Expecteds are the strings that were expecteds.
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
	} else if e.Quote {
		got = strconv.Quote(*e.Got)
	} else {
		got = *e.Got
	}

	var builder strings.Builder

	builder.WriteString("expected ")
	builder.WriteString(kind)

	var either string

	if e.Quote {
		elems := make([]string, 0, len(e.Expecteds))

		for _, elem := range e.Expecteds {
			str := strconv.Quote(elem)
			elems = append(elems, str)
		}

		either = common.EitherOrString(elems)
	} else {
		either = common.EitherOrString(e.Expecteds)
	}

	builder.WriteString(either)
	builder.WriteString(", got ")
	builder.WriteString(got)

	return builder.String()
}

// NewErrNotAsExpected creates a new ErrNotAsExpected error.
//
// Parameters:
//   - quote: Whether or not to quote the strings in the error message.
//   - kind: The kind of thing that was not as expected. This is used in the error message.
//   - got: The actual value. If nil, "nothing" is used in the error message.
//   - expecteds: The expected values. If empty, "something" is used in the error message.
//
// Returns:
//   - error: The new error. Never returns nil.
//
// Format:
//
//	"expected <kind> to be <expected>, got <got>"
//
// Where:
//   - <kind>: The kind of thing that was not as expected. This is used in the error message.
//   - <expected>: The expected values. This is used in the error message.
//   - <got>: The actual value. This is used in the error message. If nil, "nothing" is used instead.
func NewErrNotAsExpected(quote bool, kind string, got *string, expecteds ...string) error {
	return &ErrNotAsExpected{
		Quote:     quote,
		Kind:      kind,
		Expecteds: expecteds,
		Got:       got,
	}
}
