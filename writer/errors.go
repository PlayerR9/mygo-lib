package io

import "errors"

var (
	// ErrNoWriter occurs when a writer is not provided.
	//
	// Format:
	// 	"no writer was provided"
	ErrNoWriter error

	// ErrShortWrite occurs when a write operation writes fewer bytes than expected.
	//
	// Format:
	// 	"short write"
	ErrShortWrite error
)

func init() {
	ErrNoWriter = errors.New("no writer was provided")
	ErrShortWrite = errors.New("short write")
}
