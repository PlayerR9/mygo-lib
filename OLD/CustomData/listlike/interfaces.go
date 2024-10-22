package listlike

import "fmt"

// Lister is an interface that can be used by list-like data structures.
type Lister interface {
	// Size returns the number of elements in the list-like data structure.
	//
	// Returns:
	//   - int: The number of elements in the list-like data structure. Never negative.
	Size() int

	// IsEmpty checks whether the list-like data structure is empty.
	//
	// Returns:
	//   - bool: True if the list-like data structure is empty, false otherwise.
	IsEmpty() bool

	// IsNil checks whether the pointer is nil.
	//
	// Returns:
	//   - bool: True if the pointer is nil, false otherwise.
	IsNil() bool

	fmt.Stringer
}
