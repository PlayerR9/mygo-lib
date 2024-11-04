package trees

import (
	"fmt"
	"iter"
	"slices"
	"strings"

	gptr "github.com/PlayerR9/mygo-lib/pointers"
)

// Tree is a data structure that represents a tree.
type Tree[T interface {
	TreeNoder
}] interface {
	// Size returns the number of nodes in the tree.
	//
	// Returns:
	//   - int: The number of nodes in the tree.
	Size() int

	// Root returns the root node of the tree.
	//
	// Returns:
	//   - T: The root node of the tree.
	Root() T

	fmt.Stringer
}

// baseTree is a data structure that represents a tree.
type baseTree[T interface {
	Child() iter.Seq[T]
	BackwardChild() iter.Seq[T]

	TreeNoder
}] struct {
	// root is the root node of the tree.
	root T

	// leaves are the leaf nodes of the tree.
	leaves []T

	// size is the number of nodes in the tree.
	size int

	// string_fn is a function that is used to write the tree to a string.
	string_fn TreeWriterFn[T]
}

// Size implements the Tree interface.
func (t baseTree[T]) Size() int {
	return t.size
}

// String implements the Tree interface.
func (t baseTree[T]) String() string {
	var builder strings.Builder

	if t.string_fn == nil {
		t.string_fn = MakeWriteTree[T]()
	}

	_, err := t.string_fn(&builder, t.root)
	if err != nil {
		panic(err)
	}

	return builder.String()
}

// Root implements the Tree interface.
func (t baseTree[T]) Root() T {
	return t.root
}

// New creates a new Tree given the root node.
//
// Parameters:
//   - root: The root node of the tree.
//
// Returns:
//   - Tree[T]: The tree. Nil if root is nil.
func New[T interface {
	Child() iter.Seq[T]
	BackwardChild() iter.Seq[T]

	TreeNoder
}](root T) Tree[T] {
	if root.IsNil() {
		return nil
	}

	var leaves []T
	stack := []T{root}
	var size int

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		size++

		if top.IsLeaf() {
			leaves = append(leaves, top)

			continue
		}

		slice := slices.Collect(top.BackwardChild())
		gptr.RejectNils(&slice)
		stack = append(stack, slice...)
	}

	return &baseTree[T]{
		root:      root,
		leaves:    leaves,
		size:      size,
		string_fn: nil,
	}
}

// Equals checks if two trees are equal.
//
// The two trees are considered equal if they have the same number of nodes and the same structure.
// The contents of the nodes are compared with the Equals method of the node type.
//
// Parameters:
//   - other: The other tree to compare with.
//
// Returns:
//   - bool: True if the two trees are equal, false otherwise.
func Equals[T interface {
	Child() iter.Seq[T]
	Equals(other T) bool

	TreeNoder
}](tree, other Tree[T]) bool {
	if tree == nil || other == nil || tree.Size() != other.Size() {
		return false
	}

	queue := [][2]T{{tree.Root(), other.Root()}}

	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]

		if !first[0].Equals(first[1]) {
			return false
		}

		s1 := slices.Collect(first[0].Child())
		s2 := slices.Collect(first[1].Child())

		if len(s1) != len(s2) {
			return false
		}

		slice := make([][2]T, 0, len(s1))

		for i, elem := range s1 {
			slice = append(slice, [2]T{elem, s2[i]})
		}

		queue = append(queue, slice...)
	}

	return true
}
