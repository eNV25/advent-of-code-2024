package iterutil

import (
	"iter"
)

func Zip[S1 ~[]E1, S2 ~[]E2, E1 any, E2 any](s1 S1, s2 S2) iter.Seq2[E1, E2] {
	return func(yield func(E1, E2) bool) {
		for i := range min(len(s1), len(s2)) {
			if !yield(s1[i], s2[i]) {
				return
			}
		}
	}
}
