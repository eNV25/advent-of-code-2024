package sliceutil

import "cmp"

func Count[T cmp.Ordered](s []T, t T) int {
	var i int
	for _, v := range s {
		if v == t {
			i++
		}
	}
	return i
}
