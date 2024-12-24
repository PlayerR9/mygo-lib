package slices

import "github.com/PlayerR9/mygo-lib/errors"

// Builder is a builder for slices. It is only efficent for making many slices one after the other.
//
// An empty builder can be created with the `var b Builder[T]` syntax or with the
// `new(Builder[T])` constructor.
type Builder[E any] struct {
	// slice is the underlying slice being built.
	slice []E
}

// Append appends an element to the slice being built.
//
// Parameters:
//   - elem: The element to append.
//
// Returns:
//   - error: An error if the element could not be appended.
//
// Errors:
//   - common.ErrNilReceiver: If the receiver is nil.
func (b *Builder[E]) Append(elem E) error {
	if b == nil {
		return errors.ErrNilReceiver
	}

	b.slice = append(b.slice, elem)

	return nil
}

// Build builds the slice being built.
//
// Returns:
//   - []E: The slice being built.
func (b Builder[E]) Build() []E {
	if len(b.slice) == 0 {
		return nil
	}

	slice := make([]E, len(b.slice))
	copy(slice, b.slice)

	return slice
}

// Reset resets the builder to its initial state, freeing up any allocated memory
// so that it can be reused to build another slice. This is more efficient than
// creating a new builder if you need to build many slices in a row.
//
// Returns:
//   - error: An error if the builder is nil.
//
// Errors:
//   - errors.ErrNilReceiver: If the receiver is nil.
func (b *Builder[E]) Reset() error {
	if b == nil {
		return errors.ErrNilReceiver
	}

	if len(b.slice) == 0 {
		return nil
	}

	clear(b.slice)
	b.slice = nil

	return nil
}
