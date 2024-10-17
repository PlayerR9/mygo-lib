package sets

import (
	assert "github.com/PlayerR9/go-verify"
	gers "github.com/PlayerR9/mygo-lib/errors"
)

type SeenSet[K comparable] struct {
	seen map[K]struct{}
}

func NewSeenSet[K comparable]() *SeenSet[K] {
	return &SeenSet[K]{
		seen: make(map[K]struct{}),
	}
}

func (s *SeenSet[K]) See(k K) error {
	if s == nil {
		return gers.ErrNilReceiver
	}

	assert.Cond(s.seen != nil, "s.seen must not be nil")

	s.seen[k] = struct{}{}

	return nil
}

func (s SeenSet[K]) IsSeen(k K) bool {
	assert.Cond(s.seen != nil, "s.seen must not be nil")

	_, ok := s.seen[k]
	return ok
}
