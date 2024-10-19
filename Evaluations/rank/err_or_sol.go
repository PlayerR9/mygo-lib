package rank

import "errors"

// ErrOrSol is an evaluator that allows to separate errors from solutions
// whilst giving priority to solutions according to a level. Higher levels
// take precedence over lower levels.
type ErrOrSol[T any] struct {
	// errs is the list of errors.
	errs []error

	// sols is the list of solutions.
	sols []T

	// level is the current level.
	level int
}

// AddErr adds an error to the evaluator.
//
// Parameters:
//   - level: The level of the error.
//   - err: The error to add.
//
// Returns:
//   - bool: True if the receiver is not nil, false otherwise.
//
// Behaviors:
//   - If the error is nil, it is ignored.
//   - If at least a solution has been added, the error is ignored.
//   - If the error has a lower level than the current level, the error is ignored.
func (eos *ErrOrSol[T]) AddErr(level int, err error) bool {
	if err == nil {
		return true
	} else if eos == nil {
		return false
	}

	if level < eos.level || len(eos.sols) > 0 {
		return true
	}

	if level == eos.level {
		eos.errs = append(eos.errs, err)
	} else {
		eos.errs = []error{err}
		eos.level = level
	}

	return true
}

// AddSol adds a solution to the evaluator. Adding a solution will
// cause any existing errors to be ignored.
//
// Parameters:
//   - level: The level of the solution.
//   - sol: The solution to add.
//
// Returns:
//   - bool: True if the receiver is not nil, false otherwise.
//
// Behaviors:
//   - If the solution is nil, it is ignored.
//   - If at least an error has been added, the solution is ignored.
//   - If the solution has a lower level than the current level, the solution is ignored.
func (eos *ErrOrSol[T]) AddSol(level int, sol T) bool {
	if eos == nil {
		return false
	}

	if len(eos.sols) == 0 {
		if len(eos.errs) > 0 {
			for i := 0; i < len(eos.errs); i++ {
				eos.errs[i] = nil
			}

			eos.errs = eos.errs[:0]
		}

		eos.sols = []T{sol}
		eos.level = level
	} else {
		if level < eos.level {
			return true
		}

		if level == eos.level {
			eos.sols = append(eos.sols, sol)
		} else {
			eos.sols = []T{sol}
			eos.level = level
		}
	}

	return true
}

// Errs returns the list of errors.
//
// Returns:
//   - []error: The list of errors. Nil if there are no errors.
func (fos ErrOrSol[T]) Errs() []error {
	if len(fos.errs) == 0 {
		return nil
	}

	faults := make([]error, len(fos.errs))
	copy(faults, fos.errs)

	return faults
}

// Join returns an error that is the errors.Join of the errors.
//
// Returns:
//   - error: The joined error. Nil if there are no errors.
func (fos ErrOrSol[T]) Join() error {
	return errors.Join(fos.errs...)
}

// Sols returns the list of solutions.
//
// Returns:
//   - []T: The list of solutions. Nil if there are no solutions.
func (fos ErrOrSol[T]) Sols() []T {
	if len(fos.sols) == 0 {
		return nil
	}

	sols := make([]T, len(fos.sols))
	copy(sols, fos.sols)

	return sols
}

// Reset resets the evaluator to its initial state; allowing reuse. Does
// nothing if the receiver is nil.
func (fos *ErrOrSol[T]) Reset() {
	if fos == nil {
		return
	}

	if len(fos.errs) > 0 {
		for i := 0; i < len(fos.errs); i++ {
			fos.errs[i] = nil
		}

		fos.errs = fos.errs[:0]
	}

	if len(fos.sols) > 0 {
		for i := 0; i < len(fos.sols); i++ {
			fos.sols[i] = *new(T)
		}

		fos.sols = fos.sols[:0]
	}

	fos.level = 0
}
