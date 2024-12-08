package main

import (
	"fmt"
	"iter"

	"local/day7"
)

func Possibilities(input []int) iter.Seq[int] {
	return func(yield func(int) bool) {
		if len(input) == 1 {
			yield(input[0])
			return
		}
		for x := range Possibilities(input[:len(input)-1]) {
			if !yield(input[len(input)-1] * x) {
				return
			}
			if !yield(input[len(input)-1] + x) {
				return
			}
		}
	}
}

func main() {
	sum := 0
	for _, line := range day7.Input {
		n := line[0]
		operands := line[1:]
		for res := range Possibilities(operands) {
			if res == n {
				if day7.Debug {
					fmt.Println(res)
				}
				sum += n
				break
			}
		}
	}
	if day7.Debug {
		fmt.Println()
	}
	fmt.Println(sum)
}
