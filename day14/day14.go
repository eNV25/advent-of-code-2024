package day14

import (
	"bytes"
	"fmt"
	"io"
	"iter"
)

func Input() iter.Seq2[[2]int, [2]int] {
	return func(yield func([2]int, [2]int) bool) {
		r := bytes.NewReader(input)

		for {
			var p, v [2]int

			_, err := fmt.Fscanf(r, "p=%d,%d v=%d,%d\n", &p[0], &p[1], &v[0], &v[1])
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				break
			}
			if err != nil {
				panic(err)
			}

			if !yield(p, v) {
				return
			}
		}
	}
}
