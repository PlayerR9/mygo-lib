package display

import (
	"context"
	"log"

	"github.com/PlayerR9/mygo-lib/common"
)

// contextKey is the context key for the context.
type contextKey struct{}

// fromContext is a helper function that returns the underlying context from a
// context.Context object.
//
// Parameters:
//   - ctx: The context to get the underlying context from.
//
// Returns:
//   - *Context: The underlying context. Returns nil if the context is not found.
func fromContext(ctx context.Context) (*Context, error) {
	if ctx == nil {
		return nil, common.NewErrNilParam("ctx")
	}

	c, ok := ctx.Value(contextKey{}).(*Context)
	if !ok {
		return nil, common.NewErrBadParam("ctx", "must contain a *Context")
	}

	return c, nil
}

// Context is the display context.
type Context struct {
	// logger is the logger to use for the context.
	logger *log.Logger

	// cancel is the context cancel function.
	cancel context.CancelFunc
}

// NewContext creates a new context that includes a *Context value.
//
// If the parent context is nil, a new background context is created.
// The new context is associated with a new *Context and stored with a
// contextKey.
//
// Parameters:
//   - parent: The parent context to use as the base for the new context.
//   - logger: The logger to use for the context. If this is nil, the default
//     logger is used.
//
// Returns:
//   - context.Context: The new context with the *Context value.
//   - context.CancelFunc: The context cancel function.
func NewContext(parent context.Context, logger *log.Logger) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(parent)

	if logger == nil {
		logger = log.Default()
	}

	c := &Context{
		logger: logger,
		cancel: cancel,
	}

	// pc := fromContext(parent)
	// if pc != nil {
	// 	// TODO: Do something with pc
	// }

	ctx = context.WithValue(ctx, contextKey{}, c)

	return ctx, cancel
}
