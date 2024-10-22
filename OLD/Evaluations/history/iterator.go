package history

import (
	"fmt"
	"slices"

	gers "github.com/PlayerR9/mygo-lib/errors"
)

type Iterator[E Event, S Subject[E]] struct {
	init_fn       func() (S, error)
	invalid_paths []S
	stack         []*History[E]
	is_done       bool
}

// Execute returns a sequence of all possible subjects that can be obtained by executing the given initialisation
// function and then applying events to the subject. The initialisation function is called at most once per
// unique subject. The sequence is ordered, with the first element being the result of executing the initialisation
// function once, and the rest being the result of applying events to the previous element. If the initialisation
// function returns an error, it is skipped.
//
// If the subject has an error at any point, it is skipped. If the subject has an error at the end, it is yielded
// at the end of the sequence.
//
// If the sequence is exhausted, the function returns an empty sequence.
//
// Parameters:
//   - init_fn: The initialisation function to execute.
//
// Returns:
//   - iter.Seq[S]: A sequence of all possible subjects that can be obtained by executing the initialisation
//     function and then applying events to the subject. Never returns nil.
func NewIterator[E Event, S Subject[E]](init_fn func() (S, error)) (*Iterator[E, S], error) {
	if init_fn == nil {
		return nil, gers.NewBadParameter("init_fn", "not be nil")
	}

	h := NewHistory[E]()

	return &Iterator[E, S]{
		init_fn: init_fn,
		stack:   []*History[E]{h},
	}, nil
}

func (iter *Iterator[E, S]) Next() (S, error) {
	if iter == nil {
		return *new(S), ErrExhausted
	}

	if iter.is_done {
		if len(iter.invalid_paths) == 0 {
			return *new(S), ErrExhausted
		}

		a := iter.invalid_paths[0]
		iter.invalid_paths = iter.invalid_paths[1:]

		return a, nil
	}

	var subject S
	var err error

	for len(iter.stack) > 0 {
		top := iter.stack[len(iter.stack)-1]
		iter.stack = iter.stack[:len(iter.stack)-1]

		subject, err = iter.init_fn()
		if err != nil {
			return *new(S), fmt.Errorf("failed to execute initialisation function: %w", err)
		}

		other_paths, exec_err := ExecuteUntil(subject, top)

		slices.Reverse(other_paths)
		iter.stack = append(iter.stack, other_paths...)

		if exec_err == ErrErrorEncountered {
			iter.invalid_paths = append(iter.invalid_paths, subject)
		} else if exec_err != nil {
			return *new(S), fmt.Errorf("failed to execute until: %w", exec_err)
		} else {
			return subject, nil
		}
	}

	iter.is_done = true

	if len(iter.invalid_paths) == 0 {
		return *new(S), ErrExhausted
	}

	a := iter.invalid_paths[0]
	iter.invalid_paths = iter.invalid_paths[1:]

	return a, nil
}
