package mem

// Ref is a utility type for manual memory management. In essence, given a pointer `p`,
// Ref allows for `p` to be borrowed by any other function without exposing their `Free()`
// method; thus, preventing them from freeing the memory used by `p`.
type Ref struct {
	// ptr is the pointer that is being borrowed.
	ptr Freeable
}

// NewRef creates a new Ref instance that wraps a given Freeable pointer.
//
// Parameters:
//   - ptr: The Freeable pointer to be wrapped.
//
// Returns:
//   - *Ref: A new Ref instance that holds the provided pointer. Never returns nil-
//
// Panics:
//   - If the provided pointer is nil.
func NewRef(ptr Freeable) *Ref {
	if ptr == nil {
		panic("no pointer was specified")
	}

	return &Ref{
		ptr: ptr,
	}
}

// Borrow borrows the value stored by a given Ref as a given type T.
//
// Parameters:
//   - r: The Ref that holds the value to be borrowed.
//
// Returns:
//   - T: The value stored by the Ref as type T.
//
// Panics:
//   - If the Ref is nil.
//   - If the value stored by the Ref is not of type T.
func Borrow[T any](r *Ref) T {
	if r == nil {
		panic("no Ref was specified")
	}

	ptr := r.ptr
	if ptr == nil {
		panic("object is no longer valid")
	}

	val, ok := ptr.(T)
	if !ok {
		panic("reference is not of the specified type")
	}

	return val
}
