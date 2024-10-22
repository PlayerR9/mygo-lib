package sets

import (
	"fmt"
	"strings"

	assert "github.com/PlayerR9/go-verify"
	"github.com/PlayerR9/mygo-lib/common"
)

// SeenSet is a set of seen elements.
type SeenSet[K comparable] struct {
	// seen is the set of seen elements.
	seen map[K]struct{}
}

// IsEmpty implements the Set interface.
func (s SeenSet[K]) IsEmpty() bool {
	return len(s.seen) == 0
}

// Size implements the Set interface.
func (s SeenSet[K]) Size() int {
	return len(s.seen)
}

// String implements the Set interface.
func (s SeenSet[K]) String() string {
	elems := make([]string, 0, len(s.seen))

	for _, elem := range s.seen {
		elems = append(elems, fmt.Sprint(elem))
	}

	return "SeenSet[" + strings.Join(elems, ", ") + "]"
}

// NewSeenSet creates a new seen set.
//
// Returns:
//   - *SeenSet[K]: The new seen set. Never returns nil.
func NewSeenSet[K comparable]() *SeenSet[K] {
	return &SeenSet[K]{
		seen: make(map[K]struct{}),
	}
}

// See adds an element to the seen set.
//
// Parameters:
//   - k: The element to add.
//
// Returns:
//   - error: An error if the element could not be added to the set.
//
// Errors:
//   - errors.ErrNilReceiver: If the receiver is nil.
func (s *SeenSet[K]) See(k K) error {
	if s == nil {
		return common.ErrNilReceiver
	}

	assert.Cond(s.seen != nil, "s.seen must not be nil")

	s.seen[k] = struct{}{}

	return nil
}

// IsSeen checks if an element is in the seen set.
//
// Parameters:
//   - k: The element to check.
//
// Returns:
//   - bool: True if the element is in the set, false otherwise.
func (s SeenSet[K]) IsSeen(k K) bool {
	assert.Cond(s.seen != nil, "s.seen must not be nil")

	_, ok := s.seen[k]
	return ok
}
