package sets

import (
	"fmt"
)

// Set is an interface that can be used by sets.
type Set interface {
	// Size returns the number of elements in the set.
	//
	// Returns:
	//   - int: The number of elements in the set. Never negative.
	Size() int

	// IsEmpty checks whether the set is empty.
	//
	// Returns:
	//   - bool: True if the set is empty, false otherwise.
	IsEmpty() bool

	fmt.Stringer
}
