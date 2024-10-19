package history

import (
	"errors"
	"fmt"
	"iter"

	assert "github.com/PlayerR9/go-verify"
	ll "github.com/PlayerR9/mygo-lib/CustomData/listlike"
)

func Align[E Event](subject Subject[E], history *History[E]) error {
	if subject == nil {
		return fmt.Errorf("parameter (%q) must not be nil", "subject")
	} else if history == nil {
		return fmt.Errorf("parameter (%q) must not be nil", "history")
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

				ok := path.AddEvent(next)
				assert.True(ok, "path.AddEvent(next)")

				paths = append(paths, path)
			}

			alternative_paths = append(alternative_paths, paths...)
		}

		ok := history.AddEvent(nexts[0])
		assert.True(ok, "history.AddEvent(nexts[0])")

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

func Execute[E Event, S Subject[E]](init_fn func() (S, error)) iter.Seq[S] {
	if init_fn == nil {
		return func(yield func(S) bool) {}
	}

	return func(yield func(S) bool) {
		var h History[E]

		var stack ll.Stack[*History[E]]

		err := stack.Push(&h)
		assert.Err(err, "stack.Push(&h)")

		var invalid_paths []S

		for {
			top, err := stack.Pop()
			if err != nil {
				break
			}

			subject, err := init_fn()
			if err != nil {
				panic(err)
			}

			if subject.HasError() {
				invalid_paths = append(invalid_paths, subject)
				continue
			}

			other_paths, err := ExecuteUntil(subject, top)
			if err != nil {
				panic(err)
			}

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

type Runner[E Event, S Subject[E]] struct {
	init_fn func() (S, error)
}

func (r Runner[E, S]) Run() iter.Seq[S] {
	return Execute(r.init_fn)
}
