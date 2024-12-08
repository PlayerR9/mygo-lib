package common

import "github.com/PlayerR9/mygo-lib/OLD/mem"

// Typer is an interface that all types in this library must implement.
type Typer interface {
	mem.Freeable
}
