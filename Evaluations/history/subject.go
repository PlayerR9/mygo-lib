package history

// When subject has an error, then methods such as ApplyEvent must return true.
// Similarly, the NextEvents must return no events if it has an error and vice versa.
type Subject[E Event] interface {
	ApplyEvent(event E) (bool, error)
	NextEvents() ([]E, error)
	HasError() bool
}
