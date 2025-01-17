package io

// Writer is the interface that wraps the basic Write method.
type Writer interface {
	// Write writes len(p) bytes from p to the underlying data stream. It returns
	// the number of bytes written from p (0 <= n <= len(p)) and any error encountered
	// that caused the write to stop early.
	//
	// Parameters:
	//   - p: The bytes to write.
	//
	// Returns:
	//   - n: The number of bytes written from p. (0 <= n <= len(p))
	//   - err: An error if the write operation fails.
	//
	// Behaviors:
	// 	- Write must return a non-nil error if it returns n < len(p).
	// 	- Write must not modify the slice data, even temporarily.
	// 	- Implementations must not retain p.
	Write(p []byte) (int, error)
}
