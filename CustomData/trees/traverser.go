package trees

import (
	"iter"
	"slices"

	"github.com/PlayerR9/mygo-lib/common"
	gptr "github.com/PlayerR9/mygo-lib/pointers"
)

// DoFunc is a function that is called for each node in the tree.
//
// Parameters:
//   - node: The current node in the tree. Assumed to be non-nil.
//   - trav: A pointer to the traverser. Assumed to be non-nil.
//
// Returns:
//   - error: An error that indicates if the traversal should stop.
type DoFunc[T interface {
	Child() iter.Seq[T]
	BackwardChild() iter.Seq[T]

	TreeNoder
}, I any] func(node T, trav *Traversor[T, I]) error

// InfoTable is a table that stores information about each node in the tree.
type InfoTable[T interface {
	Child() iter.Seq[T]
	BackwardChild() iter.Seq[T]

	TreeNoder
}, I any] struct {
	// info is the underlying map that stores the information.
	info map[T]I
}

// Traversor is a data structure that holds information necessary for the tree
// traversal.
type Traversor[T interface {
	Child() iter.Seq[T]
	BackwardChild() iter.Seq[T]

	TreeNoder
}, I any] struct {
	// do is the function that is called for each node in the tree.
	do DoFunc[T, I]

	// info is the underlying map that stores the information.
	info map[T]I
}

// NewTraversor returns a new traverser. If no do function is provided, then
// the traverser will do nothing and return nil.
//
// Parameters:
//   - do_fn: The function that is called for each node in the tree.
//
// Returns:
//   - Traversor: A traverser.
func NewTraversor[T interface {
	Child() iter.Seq[T]
	BackwardChild() iter.Seq[T]

	TreeNoder
}, I any](do_fn DoFunc[T, I]) Traversor[T, I] {
	return Traversor[T, I]{
		do:   do_fn,
		info: make(map[T]I),
	}
}

// Set sets the information for the given node. If the node was already
// setted beforehand, then it will be overwritten.
//
// Parameters:
//   - node: The node to set the information for.
//   - info: The information to set.
//
// Returns:
//   - error: An error if the receiver is nil.
func (it *Traversor[T, I]) Set(node T, info I) error {
	if it == nil {
		return common.ErrNilReceiver
	}

	it.info[node] = info

	return nil
}

// Get retrieves the information for the given node.
//
// Parameters:
//   - node: The node to get the information for.
//
// Returns:
//   - I: The information for the given node.
//   - bool: True if the receiver is not nil and the node was found, false
//     otherwise.
func (it Traversor[T, I]) Get(node T) (I, bool) {
	if len(it.info) == 0 {
		return *new(I), false
	}

	info, ok := it.info[node]
	return info, ok
}

// DFS performs a depth-first traversal of the tree without using recursion; stopping at the
// first error encountered.
//
// Parameters:
//   - node: The node to start the traversal from.
//
// Returns:
//   - error: An error if any occurred during the traversal.
//
// Behaviors:
//   - If the node is nil or if the do function is nil, then the traversal won't
//     be performed but a nil error will be returned.
//   - Nil children will be ignored.
func (t Traversor[T, I]) DFS(node T) error {
	if node.IsNil() || t.do == nil {
		return nil
	}

	stack := []T{node}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		err := t.do(top, &t)
		if err != nil {
			return err
		}

		children := slices.Collect(top.Child())
		gptr.RejectNils(&children)

		if len(children) == 0 {
			continue
		}

		slices.Reverse(children)
		stack = append(stack, children...)
	}

	return nil
}

// BFS performs a breadth-first traversal of the tree without using recursion; stopping at the
// first error encountered.
//
// Parameters:
//   - node: The node to start the traversal from.
//
// Returns:
//   - error: An error if any occurred during the traversal.
//
// Behaviors:
//   - If the node is nil or if the do function is nil, then the traversal won't
//     be performed but a nil error will be returned.
//   - Nil children will be ignored.
func (t Traversor[T, I]) BFS(node T) error {
	if node.IsNil() || t.do == nil {
		return nil
	}

	queue := []T{node}

	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]

		err := t.do(first, &t)
		if err != nil {
			return err
		}

		children := slices.Collect(first.Child())
		gptr.RejectNils(&children)

		if len(children) == 0 {
			continue
		}

		queue = append(queue, children...)
	}

	return nil
}
