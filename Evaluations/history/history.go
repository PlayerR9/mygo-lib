package history

import (
	"iter"

	gers "github.com/PlayerR9/mygo-lib/errors"
)

type Event interface{}

type History[E Event] struct {
	timeline []E
	arrow    int
}

func (h History[E]) Copy() *History[E] {
	timeline := make([]E, len(h.timeline))
	copy(timeline, h.timeline)

	return &History[E]{
		timeline: timeline,
		arrow:    h.arrow,
	}
}

func (h *History[E]) AddEvent(event E) bool {
	if h == nil {
		return false
	}

	h.timeline = append(h.timeline, event)

	return true
}

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

func (h *History[E]) ApplyOnce(subject Subject[E]) (bool, error) {
	if h == nil {
		return false, gers.ErrNilReceiver
	} else if subject == nil {
		return false, gers.NewBadParameter("subject", "not be nil")
	}

	if h.arrow >= len(h.timeline) {
		return true, ErrHistoryDone
	}

	event := h.timeline[h.arrow]
	h.arrow++

	is_done, err := subject.ApplyEvent(event)
	if err != nil {
		return false, err
	}

	return is_done, nil
}

func (h *History[E]) Restart() {
	if h == nil {
		return
	}

	h.arrow = 0
}
