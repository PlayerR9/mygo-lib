package result

import (
	"context"
	"errors"

	"sync"

	gctx "github.com/PlayerR9/mygo-lib/Contexts"
	"github.com/PlayerR9/mygo-lib/Contexts/catcher"
	"github.com/PlayerR9/mygo-lib/Evaluation/result/internal"
	"github.com/PlayerR9/mygo-lib/common"
)

// RunFn is a function that executes a function.
type RunFn func()

// RunErrFn is a function that executes a function and returns an error.
//
// Parameters:
//   - fn: The function to execute.
//
// Returns:
//   - error: An error if the execution fails.
type RunErrFn[T Resulter] func(ctx context.Context, ch chan internal.Pair[T], elem T) error

// ExecuteBatch is a function that runs a batch of functions concurrently.
//
// Parameters:
//   - parent: The parent context.
//   - elems: The elements to create the RunErrFn from.
//   - run_fn: The RunErrFn function to execute.
//
// Returns:
//   - error: An error if any of the functions return an error.
func ExecuteBatch[T Resulter](parent context.Context, ch chan internal.Pair[T], elems []T, run_fn RunErrFn[T]) error {
	if parent == nil {
		return common.NewErrNilParam("parent")
	} else if len(elems) == 0 {
		return nil
	} else if run_fn == nil {
		return common.NewErrNilParam("make_run_fn")
	} else if ch == nil {
		return common.NewErrNilParam("ch")
	}

	// 1. Initialize the variables.
	errCh := make(chan error, len(elems))

	ctx := catcher.NewContext[error](parent)
	defer func() {
		gctx.Close(errCh)

		_ = gctx.Run(ctx, catcher.ResetCache[error]())
		// assert.Err(err, "Run(ctx, ResetCache[error]())")
	}()

	var wg sync.WaitGroup

	// 2. Make the functions.
	fns := make([]RunFn, 0, len(elems))

	for _, elem := range elems {
		fn := func() {
			defer wg.Done()

			err := run_fn(ctx, ch, elem)
			if err != nil {
				_ = gctx.Send(errCh, err)
			}
		}

		fns = append(fns, fn)
	}

	fns = fns[:len(fns):len(fns)]

	// 3. Run the functions.
	act, err := catcher.ListenCh(errCh)
	if err != nil {
		return err
	}

	err = gctx.Run(ctx, act)
	if err != nil {
		return err
	}

	wg.Add(len(fns))

	for _, fn := range fns {
		go fn()
	}

	// 4. Finalize.

	wg.Wait()
	close(errCh)

	var errs []error

	act, err = catcher.GetElems(&errs)
	if err != nil {
		return err
	}

	err = gctx.Run(ctx, catcher.Wait[error](), act)
	if err != nil {
		return err
	}

	return errors.Join(errs...)
}
