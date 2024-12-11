package day11

import (
	"fmt"
	"io"
	"strings"
)

func Input() []int {
	var s []int

	input := strings.NewReader(input)
	for {
		var n int
		_, err := fmt.Fscan(input, &n)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		s = append(s, n)
	}
	return s
}
