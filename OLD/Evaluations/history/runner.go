package history

import (
	"errors"
	"fmt"

	assert "github.com/PlayerR9/go-verify"
	"github.com/PlayerR9/mygo-lib/common"
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
	if history == nil {
		return nil
	}

	if subject == nil {
		return common.NewErrNilParam("subject")
	}

	assert.Cond(history.arrow <= len(history.timeline), "history.arrow is out of range")

	for _, event := range history.timeline[history.arrow:] {
		is_done, err := subject.ApplyEvent(event)
		if err != nil {
			return fmt.Errorf("error while applying event: %w", err)
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
		return nil, common.NewErrNilParam("subject")
	}

	if history == nil {
		return nil, common.NewErrNilParam("history")
	}

	err := Align(subject, history)
	if err != nil {
		return nil, err
	}

	var alternative_paths []*History[E]

	for {
		nexts, err := subject.NextEvents()
		if err != nil {
			return alternative_paths, err
		} else if len(nexts) == 0 {
			break
		}

		if len(nexts) > 1 {
			paths := make([]*History[E], 0, len(nexts)-1)

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
			return alternative_paths, err
		} else if is_done {
			break
		}
	}

	return alternative_paths, nil
}
