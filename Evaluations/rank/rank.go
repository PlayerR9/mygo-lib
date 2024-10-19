package rank

import (
	"slices"

	gers "github.com/PlayerR9/mygo-lib/errors"
)

type RankSol[T any] struct {
	buckets map[int][]T
	sizes   []int
}

func (r *RankSol[T]) Add(rank int, sol T) error {
	if r == nil {
		return gers.ErrNilReceiver
	}

	if r.buckets == nil {
		r.buckets = make(map[int][]T)
	}

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

func (r RankSol[T]) Build() []T {
	if len(r.sizes) == 0 {
		return nil
	}

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

func (r *RankSol[T]) Reset() {
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
