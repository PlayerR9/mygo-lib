package rank

import (
	"slices"

	assert "github.com/PlayerR9/go-verify"
	"github.com/PlayerR9/mygo-lib/common"
)

// Rank is a rank solution.
type Rank[T any] struct {
	// buckets is a map of ranks to slices of solutions.
	buckets map[int][]T

	// sizes is a sorted list of ranks.
	sizes []int
}

// New creates a new rank solution.
//
// Returns:
//   - *Rank[T]: The new rank solution. Never returns nil.
func New[T any]() *Rank[T] {
	return &Rank[T]{
		buckets: make(map[int][]T),
	}
}

// Add adds a solution to the rank solution.
//
// Parameters:
//   - rank: The rank of the solution.
//   - sol: The solution to add.
//
// Returns:
//   - error: An error if the receiver is nil.
func (r *Rank[T]) Add(rank int, sol T) error {
	if r == nil {
		return common.ErrNilReceiver
	}

	assert.Cond(r.buckets != nil, "r.buckets must not be nil")

	prev, ok := r.buckets[rank]
	if !ok {
		r.buckets[rank] = []T{sol}

		pos, ok := slices.BinarySearch(r.sizes, rank)
		if !ok {
			r.sizes = slices.Insert(r.sizes, pos, rank)
		}
	} else {
		r.buckets[rank] = append(prev, sol)
	}

	return nil
}

// Build builds the rank solution. The first element has the highest
// rank, and the last element has the lowest rank.
//
// Returns:
//   - []T: The rank solution, sorted by the rank.
func (r Rank[T]) Build() []T {
	if len(r.sizes) == 0 {
		return nil
	}

	assert.Cond(r.buckets != nil, "r.buckets must not be nil")

	max := r.sizes[len(r.sizes)-1]
	min := r.sizes[0]

	var sols []T

	for i := max; i >= min; i-- {
		bucket, ok := r.buckets[i]
		if !ok {
			continue
		}

		sols = append(sols, bucket...)
	}

	return sols
}

// Reset resets the evaluator to its initial state; allowing reuse.
func (r *Rank[T]) Reset() {
	if r == nil {
		return
	}

	if len(r.buckets) > 0 {
		for k, v := range r.buckets {
			for i := 0; i < len(v); i++ {
				v[i] = *new(T)
			}

			r.buckets[k] = nil
			delete(r.buckets, k)
		}

		r.buckets = nil
	}

	if len(r.sizes) > 0 {
		r.sizes = r.sizes[:0]
	}
}
