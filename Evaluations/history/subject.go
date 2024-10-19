package history

// Subject is a subject that can be applied to a history. In this interface, any error returned
// by the ApplyEvent and NextEvents methods are treated as panic-level error and so, they immediately
// stop the execution of the history.
//
// Therefore, the HasError method is used for non-panic error handling.
//
// Of course, when the subject's HasError method returns true, then ApplyEvent must return true as well.
// Likewise, the NextEvents must return no events when the subject's HasError method returns true and vice versa.
type Subject[E Event] interface {
	// ApplyEvent applies an event to the subject.
	//
	// Parameters:
	//   - event: The event to apply.
	//
	// Returns:
	//   - bool: True if the subject is done, false otherwise.
	//   - error: An error if the event could not be applied.
	ApplyEvent(event E) (bool, error)

	// NextEvents returns the next events in the subject.
	//
	// Returns:
	//   - []E: The next events in the subject.
	//   - error: An error if the next events could not be returned.
	NextEvents() ([]E, error)

	// HasError checks whether the subject has an error.
	//
	// Returns:
	//   - bool: True if the subject has an error, false otherwise.
	HasError() bool
}
