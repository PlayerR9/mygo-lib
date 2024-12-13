package strings

import (
	"strings"

	"github.com/PlayerR9/mygo-lib/OTHER/common"
)

// Builder is a builder for strings.
//
// An empty builder can be created with the `var b Builder` syntax or with the
// `new(Builder)` constructor.
type Builder struct {
	// builder is the underlying string builder.
	builder strings.Builder
}

// Reset implements common.Resetter.
func (b *Builder) Reset() error {
	if b == nil {
		return common.ErrNilReceiver
	}

	b.builder.Reset()

	return nil
}

// AppendString appends a string to the string being built.
//
// Parameters:
//   - s: The string to append.
//
// Returns:
//   - error: An error if the string could not be appended.
//
// Errors:
//   - common.ErrNilReceiver: If the receiver is nil.
func (b *Builder) AppendString(s string) error {
	if b == nil {
		return common.ErrNilReceiver
	}

	if s == "" {
		return nil
	}

	data := []byte(s)

	_, _ = b.builder.Write(data)

	return nil
}

// Build builds the string being built.
//
// Returns:
//   - string: The string being built.
//
// Note: This function does not reset the builder.
func (b Builder) Build() string {
	str := b.builder.String()
	return str
}
