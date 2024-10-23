package history

import (
	"iter"

	"github.com/PlayerR9/mygo-lib/common"
)

// Event is an event that can be applied to a subject.
type Event interface{}

// History is a history of events.
type History[E Event] struct {
	// timeline is the history of events.
	timeline []E

	// arrow is the index of the next event to be applied.
	arrow int
}

// NewHistory creates a new history.
//
// Returns:
//   - *History: The new history. Never returns nil.
//
// This function is just a convenience function that does the same as:
//
//	var h History[E]
func NewHistory[E Event]() *History[E] {
	return &History[E]{
		timeline: make([]E, 0),
		arrow:    0,
	}
}

// Copy creates a copy of the history.
//
// Returns:
//   - *History: The copy. Never returns nil.
func (h History[E]) Copy() *History[E] {
	timeline := make([]E, len(h.timeline))
	copy(timeline, h.timeline)

	return &History[E]{
		timeline: timeline,
		arrow:    h.arrow,
	}
}

// AddEvent adds an event to the history.
//
// Returns:
//   - error: An error if the receiver is nil.
func (h *History[E]) AddEvent(event E) error {
	if h == nil {
		return common.ErrNilReceiver
	}

	h.timeline = append(h.timeline, event)

	return nil
}

// Event iterates over the events in the history. Once iteration stops, the history
// is advanced to where it stopped.
//
// Returns:
//   - iter.Seq[E]: The events in the history. Never returns nil.
func (h *History[E]) Event() iter.Seq[E] {
	if h == nil {
		return func(yield func(E) bool) {}
	}

	return func(yield func(E) bool) {
		for i := h.arrow; i < len(h.timeline); i++ {
			if !yield(h.timeline[i]) {
				h.arrow = i
				break
			}
		}

		h.arrow = len(h.timeline)
	}
}

// ApplyOnce applies the next event in the history to the subject.
//
// Parameters:
//   - subject: The subject to apply the next event to.
//
// Returns:
//   - bool: True if the subject is done, false otherwise.
//   - error: An error if the receiver or subject is nil.
//
// Errors:
//   - errors.ErrNilReceiver: The receiver is nil.
//   - ErrHistoryDone: The history is done.
//   - any other error: if the subject is nil or any error returned by the subject.ApplyEvent() method.
func (h *History[E]) ApplyOnce(subject Subject[E]) (bool, error) {
	if h == nil {
		return false, common.ErrNilReceiver
	} else if subject == nil {
		return false, common.NewErrNilParam("subject")
	}

	if h.arrow >= len(h.timeline) {
		return true, ErrHistoryDone
	}

	event := h.timeline[h.arrow]
	h.arrow++

	is_done, err := subject.ApplyEvent(event)
	if err != nil {
		return true, err
	}

	return is_done, nil
}

// Restart restarts the history for reuse.
func (h *History[E]) Restart() {
	if h == nil {
		return
	}

	h.arrow = 0
}
