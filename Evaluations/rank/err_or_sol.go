package rank

import (
	"errors"

	assert "github.com/PlayerR9/go-verify"
	gers "github.com/PlayerR9/mygo-lib/errors"
)

// ErrOrSol is an evaluator that allows to separate errors from solutions
// whilst giving priority to solutions according to a level. Higher levels
// take precedence over lower levels.
type ErrOrSol[T any] struct {
	errs *Rank[error]
	sols *Rank[T]
}

// NewErrOrSol returns a new ErrOrSol.
//
// Returns:
//   - *ErrOrSol[T]: The new evaluator. Never returns nil.
func NewErrOrSol[T any]() *ErrOrSol[T] {
	return &ErrOrSol[T]{
		errs: New[error](),
		sols: nil,
	}
}

// AddErr adds an error to the evaluator.
//
// Parameters:
//   - rank: The level of the error.
//   - e: The error to add.
//
// Returns:
//   - bool: True if the error could be added, false otherwise.
//
// Behaviors:
//   - If the error is nil, it is ignored.
//   - If at least a solution has been added, the error is ignored.
func (eos ErrOrSol[T]) AddErr(rank int, e error) bool {
	if e == nil {
		return true
	} else if eos.errs == nil {
		return false
	}

	err := eos.errs.Add(rank, e)
	assert.Err(err, "eos.errs.Add(rank, e)")

	return true
}

// AddSol adds a solution to the evaluator. Adding a solution will
// cause any existing errors to be ignored.
//
// Parameters:
//   - rank: The level of the solution.
//   - sol: The solution to add.
//
// Returns:
//   - error: An error if the receiver is nil.
func (eos *ErrOrSol[T]) AddSol(rank int, sol T) error {
	if eos == nil {
		return gers.ErrNilReceiver
	}

	if eos.sols == nil {
		eos.sols = New[T]()
	} else if eos.errs != nil {
		eos.errs = nil
	}

	err := eos.sols.Add(rank, sol)
	assert.Err(err, "eos.sols.Add(level, sol)")

	return nil
}

// Errs returns the list of errors.
//
// Returns:
//   - []error: The list of errors. Nil if there are no errors.
func (eos ErrOrSol[T]) Errs() []error {
	if eos.errs == nil {
		return nil
	}

	errs := eos.errs.Build()
	return errs
}

// Join returns an error that is the errors.Join of the errors.
//
// Returns:
//   - error: The joined error. Nil if there are no errors.
func (eos ErrOrSol[T]) Join() error {
	if eos.errs == nil {
		return nil
	}

	errs := eos.errs.Build()
	err := errors.Join(errs...)

	return err
}

// Sols returns the list of solutions.
//
// Returns:
//   - []T: The list of solutions. Nil if there are no solutions.
func (eos ErrOrSol[T]) Sols() []T {
	if eos.sols == nil {
		return nil
	}

	sols := eos.sols.Build()
	return sols
}

// Reset resets the evaluator to its initial state; allowing reuse.
func (eos *ErrOrSol[T]) Reset() {
	if eos == nil {
		return
	}

	eos.errs.Reset()
	eos.sols.Reset()

	eos.errs = New[error]()
	eos.sols = nil
}
