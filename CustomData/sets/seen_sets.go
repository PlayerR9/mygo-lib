package sets

type SeenSet[K comparable] struct {
	seen map[K]struct{}
}

func NewSeenSet[K comparable]() *SeenSet[K] {
	return &SeenSet[K]{
		seen: make(map[K]struct{}),
	}
}
