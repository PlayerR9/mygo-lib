package mem

// Ref is a utility type for manual memory management. In essence, given a pointer `p`,
// Ref allows for `p` to be borrowed by any other function without exposing their `Free()`
// method; thus, preventing them from freeing the memory used by `p`.
type Ref struct {
	// free is the function that will be called when the Ref is freed.
	free func()

	// ptr is the pointer that is being borrowed.
	ptr any
}

// NewRef creates a new Ref instance that wraps a given pointer.
//
// Parameters:
//   - ptr: The Freeable pointer to be wrapped.
//   - free: The function that will be called when the Ref is freed.
//
// Returns:
//   - *Ref: A new Ref instance that holds the provided pointer. Never returns nil.
//
// Panics:
// - ErrBadParam: If the provided pointer is nil or the provided free function is nil.
//
// Example:
//
//	type MyType struct{
//		// Your fields go here...
//	}
//
//	func (mt *MyType) free() {
//		if mt == nil {
//			return
//		}
//
//		// Handle the cleanup procedure here....
//	}
//
//	func NewMyType(/* */) *mem.Ref {
//		mt := // Construct here your type...
//
//		ref := NewRef(mt, mt.free)
//		return ref
//	}
func NewRef(ptr any, free func()) *Ref {
	if ptr == nil {
		panic(NewErrNilParam("ptr"))
	} else if free == nil {
		panic(NewErrNilParam("free"))
	}

	return &Ref{
		ptr:  ptr,
		free: free,
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
func Borrow[T any](ref *Ref) T {
	if ref == nil {
		panic(NewErrNilParam("ref"))
	}

	if ref.free == nil {
		panic(NewErrInvalidObject("Borrow"))
	}

	val, ok := ref.ptr.(T)
	if !ok {
		panic(NewErrInvalidType(ref.ptr, *new(T)))
	}

	return val
}
