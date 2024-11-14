package mem

// Ref is a utility type for manual memory management. In essence, given a pointer `p`,
// Ref allows for `p` to be borrowed by any other function without exposing their `Free()`
// method; thus, preventing them from freeing the memory used by `p`.
type Ref[T any] struct {
	// free is the function that will be called when the Ref is freed.
	free func() error

	// ptr is the pointer that is being borrowed.
	ptr T
}

// Free implements Freeable.
func (ref *Ref[T]) Free() error {
	if ref == nil {
		return ErrNilReceiver
	} else if ref.free == nil {
		// Method `Free()` was called.
		return NewErrInvalidObject("Free")
	}

	err := ref.free()
	if err != nil {
		return err
	}

	ref.free = nil
	ref.ptr = *new(T)

	return nil
}

// New creates a new instance of Ref that allows to manually deallocate a pointer.
//
// Parameters:
//   - ptr: The pointer managed by Ref.
//   - freeFn: The function that handles the deallocation of the pointer.
//
// Returns:
//   - *Ref: A new instance of Ref with the given pointer and free function.
//
// This function returns a nil reference iff the provided freeFn function is nil.
//
// Example:
//
//	type MyType struct{
//		// Your fields go here...
//	}
//
//	func (mt *MyType) free() error {
//		if mt == nil {
//			return mem.ErrNilReceiver
//		}
//
//		// Handle the cleanup procedure here....
//
//		return nil
//	}
//
//	func NewMyType(/* */) *mem.Ref[*MyType] {
//		mt := // Construct here your type...
//
//		ref := mem.New[*MyType](mt, mt.free)
//		return ref
//	}
func New[T any](ptr T, freeFn func() error) *Ref[T] {
	if freeFn == nil {
		return nil
	}

	return &Ref[T]{
		ptr:  ptr,
		free: freeFn,
	}
}

// Borrow returns the pointer to the value.
//
// Returns:
//   - T: The pointer to the value.
//   - error: An error of type *ErrInvalidObject if the Ref is already freed.
func (ref Ref[T]) Borrow() (T, error) {
	if ref.free == nil {
		return *new(T), NewErrInvalidObject("Borrow()")
	}

	return ref.ptr, nil
}

// MustBorrow returns the pointer to the value.
//
// Panics:
//   - *ErrInvalidObject: If the Ref is already freed.
func (ref Ref[T]) MustBorrow() T {
	if ref.free == nil {
		panic(NewErrInvalidObject("Borrow()"))
	}

	return ref.ptr
}
