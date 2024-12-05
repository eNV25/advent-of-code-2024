package sliceutil

func Count[S ~[]E, E comparable](s S, e E) int {
	var i int
	for _, v := range s {
		if v == e {
			i++
		}
	}
	return i
}
