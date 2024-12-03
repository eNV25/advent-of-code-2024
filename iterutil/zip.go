package iterutil

import (
	"iter"
)

func Zip[T, U any](t []T, u []U) iter.Seq2[T, U] {
	return func(yield func(T, U) bool) {
		for i := range min(len(t), len(u)) {
			if !yield(t[i], u[i]) {
				return
			}
		}
	}
}
