package runes

import (
	"github.com/PlayerR9/mygo-lib/runes/internal"
)

// trimT contains the methods for trimming characters from a slice of runes.
type trimT struct{}

// TRIM is the namespace for Trim methods.
var TRIM trimT

// Beginning trims runes from the beginning of a slice of runes based on a predicate function.
//
// Parameters:
//   - s: The slice of runes to trim.
//   - isTrim: A function that determines if a rune should be trimmed.
//
// Returns:
//   - uint: The number of characters trimmed from the beginning.
//
// Panics:
//   - ErrNoPredicate: If the predicate function is nil.
func (trimT) Beginning(s *[]rune, isTrim func(rune) bool) uint {
	if isTrim == nil {
		panic(ErrNoPredicate)
	}

	if s == nil || len(*s) == 0 {
		return 0
	}

	n := internal.TrimLeft(s, isTrim)
	return n
}

// End trims runes from the end of a slice of runes based on a predicate function.
//
// Parameters:
//   - s: The slice of runes to trim.
//   - isTrim: A function that determines if a rune should be trimmed.
//
// Returns:
//   - uint: The number of characters trimmed from the end.
//
// Panics:
//   - ErrNoPredicate: If the predicate function is nil.
func (trimT) End(s *[]rune, isTrim func(rune) bool) uint {
	if isTrim == nil {
		panic(ErrNoPredicate)
	}

	if s == nil || len(*s) == 0 {
		return 0
	}

	n := internal.TrimRight(s, isTrim)
	return n
}

// Outer trims runes from the beginning and end of a slice of runes based on a predicate function.
//
// Parameters:
//   - s: The slice of runes to trim.
//   - isTrim: A function that determines if a rune should be trimmed.
//
// Returns:
//   - uint: The number of characters trimmed from the beginning and end.
//
// Panics:
//   - ErrNoPredicate: If the predicate function is nil.
func (trimT) Outer(s *[]rune, isTrim func(rune) bool) uint {
	if isTrim == nil {
		panic(ErrNoPredicate)
	}

	if s == nil || len(*s) == 0 {
		return 0
	}

	n := internal.Trim(s, isTrim)
	return n
}
