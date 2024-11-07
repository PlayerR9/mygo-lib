package math

// UintPow calculates base to the power exp, where base and exp are unsigned
// integers. It does this efficiently by using exponentiation by squaring.
//
// Parameters:
//   - base: The base of the exponentiation.
//   - exp: The exponent.
//
// Returns:
//   - uint: The result of the exponentiation.
//   - error: An error of type ErrInvalidBase if the base is zero.
//
// Example:
//
//	UintPow(2, 8) // Returns 256
func UintPow(base, exp uint) (uint, error) {
	if base == 0 {
		return 0, ErrInvalidBase
	} else if exp == 0 || base == 1 {
		return 1, nil
	}

	res := uint(1)

	for exp > 0 {
		if exp&1 == 1 {
			res *= base
		}

		exp >>= 1
		base *= base
	}

	return res, nil
}

// UintPowSlice calculates the slice of all powers of the given base up to the maximum exponent.
//
// Parameters:
//   - base: The base of the exponentiation.
//   - max_exp: The maximum exponent.
//
// Returns:
//   - []uint: A slice of all powers of base from 0 to max_exp.
//   - error: An error of type ErrInvalidBase if the base is zero.
//
// Example:
//
//	UintPowSlice(2, 8) // Returns [1, 2, 4, 8, 16, 32, 64, 128, 256]
func UintPowSlice(base uint, max_exp uint) ([]uint, error) {
	if base == 0 {
		return nil, ErrInvalidBase
	} else if max_exp == 0 {
		return []uint{1}, nil
	}

	slice := make([]uint, max_exp+1)
	slice[0] = 1

	prev := slice[0]
	for exp := uint(1); exp <= max_exp; exp++ {
		prev *= base
		slice[exp] = prev
	}

	return slice, nil
}
