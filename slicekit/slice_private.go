package slicekit

func sliceCopy[T any](s []T) []T {
	slice := make([]T, 0, cap(s))

	for _, v := range s {
		slice = append(slice, v)
	}

	return slice
}
