package day1

import (
	"fmt"
	"io"
)

func Input() (l1 []int, l2 []int) {
	for {
		var v1, v2 int

		_, err := fmt.Scan(&v1, &v2)
		if err == io.EOF {
			return
		}
		if err != nil {
			panic(err)
		}

		l1 = append(l1, v1)
		l2 = append(l2, v2)
	}
}
