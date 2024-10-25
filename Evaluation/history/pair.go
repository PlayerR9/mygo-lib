package history

// Pair is a pair of a solution and a history.
type Pair[E any, S Subject[E]] struct {
	// Solution is the solution.
	Solution S

	// History is the history.
	History []E
}

// NewPair creates a new pair.
//
// Parameters:
//   - solution: The solution.
//   - history: The history.
//
// Returns:
//   - Pair[E, S]: The new pair.
func NewPair[E any, S Subject[E]](solution S, history []E) Pair[E, S] {
	return Pair[E, S]{
		Solution: solution,
		History:  history,
	}
}
