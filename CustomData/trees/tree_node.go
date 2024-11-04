package trees

import (
	"fmt"
)

// TreeNoder is an interface that must be implemented by all tree nodes.
type TreeNoder interface {
	comparable

	// IsNil checks whether the node is nil.
	//
	// Returns:
	//   - bool: True if the node is nil, false otherwise.
	IsNil() bool

	// IsLeaf checks whether the node is a leaf node. A nil node is
	// considered a leaf node.
	//
	// Returns:
	//   - bool: True if the node is a leaf node, false otherwise.
	IsLeaf() bool

	fmt.Stringer
}
