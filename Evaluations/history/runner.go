package history

import (
	"errors"
	"iter"

	assert "github.com/PlayerR9/go-verify"
	ll "github.com/PlayerR9/mygo-lib/CustomData/listlike"
	gers "github.com/PlayerR9/mygo-lib/errors"
)

// Align aligns a subject with a history. Does nothing if the history is already aligned or nil.
//
// Parameters:
//   - subject: The subject to align with the history.
//   - history: The history to align with the subject.
//
// Returns:
//   - error: An error if the subject is nil, or if the subject got an error or is done before the
//     history could be aligned.
func Align[E Event](subject Subject[E], history *History[E]) error {
	if history == nil || history.arrow == len(history.timeline) {
		return nil
	}

	if subject == nil {
		return gers.NewBadParameter("subject", "not be nil")
	}

	for _, event := range history.timeline[history.arrow:] {
		is_done, err := subject.ApplyEvent(event)
		if err != nil {
			return err
		}

		if subject.HasError() {
			return errors.New("subject got an error before the history could be aligned")
		} else if is_done {
			return errors.New("subject is done before the history could be aligned")
		}
	}

	history.arrow = len(history.timeline)

	return nil
}

// ExecuteUntil executes the history until it is done or an event is found that
// was not expected by the subject.
//
// Parameters:
//   - subject: The subject to execute with the history.
//   - history: The history to execute with the subject.
//
// Returns:
//   - []*History[E]: A list of alternative paths that were found. Never returns nil.
//   - error: An error if the subject is nil, or if the subject got an error or is done before the
//     history could be aligned.
func ExecuteUntil[E Event](subject Subject[E], history *History[E]) ([]*History[E], error) {
	if subject == nil {
		return nil, nil
	}

	err := Align(subject, history)
	if err != nil {
		return nil, err
	}

	var alternative_paths []*History[E]

	for {
		nexts, err := subject.NextEvents()
		if err != nil {
			return nil, err
		}

		if len(nexts) == 0 {
			break
		}

		if len(nexts) > 1 {
			paths := make([]*History[E], 0, len(nexts))

			for _, next := range nexts[1:] {
				path := history.Copy()
				path.Restart()

				err := path.AddEvent(next)
				assert.Err(err, "path.AddEvent(next)")

				paths = append(paths, path)
			}

			alternative_paths = append(alternative_paths, paths...)
		}

		err = history.AddEvent(nexts[0])
		assert.Err(err, "history.AddEvent(nexts[0])")

		is_done, err := history.ApplyOnce(subject)
		if err != nil {
			return nil, err
		}

		if is_done {
			break
		}
	}

	return alternative_paths, nil
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
func Execute[E Event, S Subject[E]](init_fn func() (S, error)) iter.Seq[S] {
	if init_fn == nil {
		return func(yield func(S) bool) {}
	}

	return func(yield func(S) bool) {
		h := NewHistory[E]()

		stack := ll.NewStackFromSlice([]*History[E]{h})

		var invalid_paths []S

		for {
			top, err := stack.Pop()
			if err != nil {
				break
			}

			subject, err := init_fn()
			assert.Err(err, "init_fn()")

			if subject.HasError() {
				invalid_paths = append(invalid_paths, subject)
				continue
			}

			other_paths, err := ExecuteUntil(subject, top)
			assert.Err(err, "ExecuteUntil(subject, top)")

			err = stack.PushMany(other_paths)
			assert.Err(err, "stack.PushMany(other_paths)")

			if subject.HasError() {
				invalid_paths = append(invalid_paths, subject)
			} else if !yield(subject) {
				return
			}
		}

		for _, a := range invalid_paths {
			if !yield(a) {
				return
			}
		}
	}
}

// Runner executes the subject until the history is done.
type Runner[E Event, S Subject[E]] struct {
	init_fn func() (S, error)
}

// New creates a new runner.
//
// Parameters:
//   - init_fn: The function to initialize the subject.
//
// Returns:
//   - *Runner[E, S]: The runner.
//   - error: An error if the init_fn is nil.
func New[E Event, S Subject[E]](init_fn func() (S, error)) (*Runner[E, S], error) {
	if init_fn == nil {
		return nil, gers.NewBadParameter("init_fn", "not be nil")
	}

	return &Runner[E, S]{
		init_fn: init_fn,
	}, nil
}

// Run runs the runner. It returns a sequence of all the subjects that the
// runner visits. The runner may visit a subject multiple times, but it
// will be returned in the sequence only once. The sequence is ordered
// such that all the valid subjects come first, followed by all the
// invalid subjects.
//
// Returns:
//   - iter.Seq[S]: The subjects. Never returns nil.
func (r Runner[E, S]) Run() iter.Seq[S] {
	return Execute(r.init_fn)
}
