package slices

import (
	gers "github.com/PlayerR9/mygo-lib/errors"
)

// Builder is a slice builder.
type Builder[T any] struct {
	// slice is the slice being built.
	slice []T
}

// Append appends an element to the slice being built.
//
// Parameters:
//   - elem: The element to append.
//
// Returns:
//   - error: An error if the receiver is nil.
func (b *Builder[T]) Append(elem T) error {
	if b == nil {
		return gers.ErrNilReceiver
	}

	b.slice = append(b.slice, elem)

	return nil
}

// Build builds the slice being built.
//
// Returns:
//   - []T: The slice being built.
func (b Builder[T]) Build() []T {
	if len(b.slice) == 0 {
		return nil
	}

	slice := make([]T, len(b.slice), len(b.slice))
	copy(slice, b.slice)

	return slice
}

// Reset resets the builder for reuse.
func (b *Builder[T]) Reset() {
	if b == nil {
		return
	}

	if len(b.slice) > 0 {
		for i := 0; i < len(b.slice); i++ {
			b.slice[i] = *new(T)
		}

		b.slice = b.slice[:0]
	}
}
