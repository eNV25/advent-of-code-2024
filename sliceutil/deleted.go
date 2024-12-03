package sliceutil

import "slices"

func Deleted[S ~[]E, E any](s S, i, j int) (r S) {
	return slices.Delete(slices.Clone(s), i, j)
}
