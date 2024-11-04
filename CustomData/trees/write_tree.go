package trees

import (
	"io"
	"iter"

	"github.com/PlayerR9/mygo-lib/trees/internal"
)

var (
	nonterminal_data, terminal_data, indent_pipe_data, indent_empty_data []byte
	nonterminal_size, terminal_size                                      int
)

func init() {
	nonterminal_data = []byte("├── ")
	nonterminal_size = len(nonterminal_data)

	terminal_data = []byte("└── ")
	terminal_size = len(terminal_data)

	indent_pipe_data = []byte("│   ")

	indent_empty_data = []byte("    ")
}

type TreeWriterFn[T interface {
	Child() iter.Seq[T]
	BackwardChild() iter.Seq[T]

	TreeNoder
}] func(w io.Writer, root T) (int, error)

// MakeWriteTree returns a string representation of the given token tree
// in a form that is easier to read and understand.
//
// Parameters:
//   - root: The root of the token tree to print.
//
// Returns:
//   - string: A string representation of the token tree.
//
// If root is nil, an empty string is returned.
//
// WARNING: This function is recursive.
func MakeWriteTree[T interface {
	Child() iter.Seq[T]
	BackwardChild() iter.Seq[T]

	TreeNoder
}]() TreeWriterFn[T] {
	fn := func(node T, trav *Traversor[T, *internal.MakeWTInfo]) error {
		info, _ := trav.Get(node)

		if !info.IsFirst() {
			data := append([]byte("\n"), info.Indent()...)

			expected_size := len(data)

			if info.IsLast() {
				data = append(data, terminal_data...)
				expected_size += terminal_size
			} else {
				data = append(data, nonterminal_data...)
				expected_size += nonterminal_size
			}

			err := info.WriteData(data, expected_size)
			if err != nil {
				return err
			}
		}

		err := info.WriteString(node.String())
		if err != nil {
			return err
		}

		seq := node.Child()
		if seq == nil {
			return nil
		}

		var last_info *internal.MakeWTInfo
		var new_indent []byte

		if info.IsFirst() {
			new_indent = nil
		} else if info.IsLast() {
			new_indent = append([]byte(info.Indent()), indent_empty_data...)
		} else {
			new_indent = append([]byte(info.Indent()), indent_pipe_data...)
		}

		for child := range seq {
			new_info := info.New(new_indent)

			last_info = new_info

			_ = trav.Set(child, new_info)
		}

		if last_info != nil {
			last_info.SetIsLast(true)
		}

		return nil
	}

	trav := NewTraversor(fn)

	return func(w io.Writer, root T) (int, error) {
		info := internal.NewMakeWTInfo(w)

		_ = trav.Set(root, info)

		err := trav.DFS(root)
		return info.BytesWritten(), err
	}
}
