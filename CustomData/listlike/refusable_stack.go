package listlike

import (
	"fmt"

	assert "github.com/PlayerR9/go-verify"
)

type RefusableStack[T any] struct {
	elems      []T
	stack_size int
	stack_end  int
}

func (rs RefusableStack[T]) Validate() error {
	if rs.stack_size != len(rs.elems) {
		return fmt.Errorf("%q mismatch", "stack_size")
	}

	if rs.stack_end < 0 || rs.stack_end > rs.stack_size {
		return fmt.Errorf("%q is not in [%d, %d]", "stack_end", 0, rs.stack_size)
	}

	return nil
}

func (rs RefusableStack[T]) Size() int {
	err := rs.Validate()
	assert.Err(err, "rs.Validate()")

	return rs.stack_size
}

func (rs *RefusableStack[T]) Push(elem T) bool {
	if rs == nil {
		return false
	}

	err := rs.Validate()
	assert.Err(err, "rs.Validate()")

	assert.Cond(rs.stack_end == rs.stack_size, "cannot push when elements are not accepted or refused beforehand")

	rs.elems = append(rs.elems, elem)
	rs.stack_size++
	rs.stack_end++

	return true
}

func (rs *RefusableStack[T]) Pop() (T, bool) {
	zero := *new(T)

	if rs == nil {
		return zero, false
	}

	err := rs.Validate()
	assert.Err(err, "rs.Validate()")

	if rs.stack_end == 0 {
		return zero, false
	}

	rs.stack_end--

	return rs.elems[rs.stack_end], true
}

func (rs *RefusableStack[T]) Accept() {
	if rs == nil {
		return
	}

	err := rs.Validate()
	assert.Err(err, "rs.Validate()")

	rs.elems = rs.elems[:rs.stack_end]
	rs.stack_size = rs.stack_end
}

func (rs *RefusableStack[T]) Refuse() {
	if rs == nil {
		return
	}

	err := rs.Validate()
	assert.Err(err, "rs.Validate()")

	rs.stack_end = rs.stack_size
}

func (rs RefusableStack[T]) Popped() []T {
	err := rs.Validate()
	assert.Err(err, "rs.Validate()")

	slice := make([]T, 0, rs.stack_size-rs.stack_end)

	for _, elem := range rs.elems[rs.stack_end:] {
		slice = append(slice, elem)
	}

	return slice
}
