package runes

/*
// ErrNotAsExpected is an error that is returned when an element is not as expected.
type ErrNotAsExpected struct {
	// Kind is the kind of the expected elements.
	Kind string

	// Expecteds is the expected elements.
	Expecteds []rune

	// Got is the actual element.
	Got *rune
}

// Error implements the error interface.
func (e ErrNotAsExpected) Error() string {
	var expected string

	if len(e.Expecteds) == 0 {
		expected = "something"
	} else {
		elems := make([]string, 0, len(e.Expecteds))

		for _, elem := range e.Expecteds {
			elems = append(elems, strconv.QuoteRune(elem))
		}

		expected = EitherOrString(elems)
	}

	var got string

	if e.Got == nil {
		got = "nothing"
	} else {
		got = strconv.QuoteRune(*e.Got)
	}

	var builder strings.Builder

	builder.WriteString("expected ")

	if e.Kind != "" {
		builder.WriteString(e.Kind)
		builder.WriteString(" to be ")
	}

	builder.WriteString(expected)
	builder.WriteString(", got ")
	builder.WriteString(got)

	return builder.String()
}

// NewErrNotAsExpected returns a new ErrNotAsExpected error.
//
// Parameters:
//   - quote: Indicates whether the expected elements should be quoted.
//   - kind: The kind of the expected elements. If empty, it will be omitted.
//   - got: The actual element.
//   - expecteds: The expected elements.
//
// Returns:
//   - error: The new error. Never returns nil.
//
// Format:
//
//	expected <kind> to be <expecteds>, got <got>
func NewErrNotAsExpected(kind string, got *rune, expecteds ...rune) error {
	return &ErrNotAsExpected{
		Kind:      kind,
		Expecteds: expecteds,
		Got:       got,
	}
}

// ErrAfter is an error that is returned when an error occurs after an element.
type ErrAfter struct {
	// Previous is the element before the one that caused the error.
	Previous rune

	// Inner is the inner error.
	Inner error
}

// Error implements the error interface.
func (e ErrAfter) Error() string {
	return "after " + string(e.Previous) + ": " + ErrMsgOf(e.Inner)
}

// NewErrAfter returns a new ErrAfter error.
//
// Parameters:
//   - previous: The element before the one that caused the error.
//   - inner: The inner error.
//
// Returns:
//   - error: The new error. Never returns nil.
//
// Format:
//
//	"after <previous>: <inner>"
//
// Where:
//   - <previous>: The element before the one that caused the error. If empty, "something" is used.
//   - <inner>: The inner error. If nil, the default error message is used.
func NewErrAfter(previous rune, inner error) error {
	return &ErrAfter{
		Previous: previous,
		Inner:    inner,
	}
}

// Unwrap returns the inner error.
//
// Returns:
//   - error: The inner error.
func (e ErrAfter) Unwrap() error {
	return e.Inner
}
*/
