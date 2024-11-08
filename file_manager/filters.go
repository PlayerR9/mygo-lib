package file_manager

import (
	"io/fs"

	"github.com/PlayerR9/mygo-lib/common"
)

// RejectNilDirEntry filters out nil elements from a slice of fs.DirEntry pointers,
// modifying the slice in-place. If the slice is nil or empty, it does nothing.
//
// Parameters:
//   - slice: A pointer to a slice of fs.DirEntry pointers. The slice will be
//     modified to remove any nil elements. If all elements are nil, the slice
//     will be set to nil.
//
// Returns:
//   - uint: The number of elements removed from the slice.
//
// Panics:
//   - common.ErrBadParam: If the entries parameter is nil.
func RejectNilDirEntry(slice *[]fs.DirEntry) uint {
	if slice == nil {
		panic(common.NewErrNilParam("slice"))
	}

	lenSlice := uint(len(*slice))
	if lenSlice == 0 {
		return 0
	}

	var top uint

	for _, entry := range *slice {
		if entry == nil {
			continue
		}

		(*slice)[top] = entry
		top++
	}

	if top == 0 {
		clear(*slice)
		*slice = nil

		return lenSlice
	}

	clear((*slice)[top:])
	*slice = (*slice)[:top]

	return lenSlice - top
}

// RejectDir filters out nil elements and directory entries from a slice of fs.DirEntry
// pointers, modifying the slice in-place. If the slice is nil or empty, it does nothing.
//
// Parameters:
//   - slice: A pointer to a slice of fs.DirEntry pointers. The slice will be
//     modified to remove any nil or directory elements. If all elements are filtered out,
//     the slice will be set to nil.
//
// Returns:
//   - uint: The number of elements removed from the slice.
//
// Panics:
//   - common.ErrBadParam: If the entries parameter is nil.
func RejectDir(slice *[]fs.DirEntry) uint {
	if slice == nil {
		panic(common.NewErrNilParam("entries"))
	}

	lenSlice := uint(len(*slice))
	if lenSlice == 0 {
		return 0
	}

	var top uint

	for _, entry := range *slice {
		if entry == nil || entry.IsDir() {
			continue
		}

		(*slice)[top] = entry
		top++
	}

	if top == 0 {
		clear(*slice)
		*slice = nil

		return lenSlice
	}

	clear((*slice)[top:])
	*slice = (*slice)[:top]

	return lenSlice - top
}
