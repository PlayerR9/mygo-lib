package internal

func FirstIndexOf[S ~[]E, E comparable](slice S, target E) (uint, bool) {
	if len(slice) == 0 {
		return 0, false
	}

	for i, v := range slice {
		if v == target {
			return uint(i), true
		}
	}

	return 0, false
}
