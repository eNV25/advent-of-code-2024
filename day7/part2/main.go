package main

import (
	"fmt"
	"iter"

	"local/day7"
)

func concat(x, y int) (z int) {
	z = x
	for i := 1; i <= y; i *= 10 {
		z *= 10
	}
	z += y
	if day7.Debug {
		fmt.Printf("concat(%d, %d) = %d\n", x, y, z)
	}
	return
}

func mul(x, y int) (z int) {
	z = x * y
	if day7.Debug {
		fmt.Printf("mul(%d, %d) = %d\n", x, y, z)
	}
	return
}

func sum(x, y int) (z int) {
	z = x + y
	if day7.Debug {
		fmt.Printf("sum(%d, %d) = %d\n", x, y, z)
	}
	return
}

func Possible(input []int) iter.Seq[int] {
	return func(yield func(int) bool) {
		if len(input) == 1 {
			yield(input[0])
			return
		}
		for x := range Possible(input[:len(input)-1]) {
			if !yield(concat(x, input[len(input)-1])) {
				return
			}
			if !yield(mul(x, input[len(input)-1])) {
				return
			}
			if !yield(sum(x, input[len(input)-1])) {
				return
			}
		}
	}
}

func main() {
	sum := 0
	for _, line := range day7.Input {
		n := line[0]
		oprnds := line[1:]
		for pssble := range Possible(oprnds) {
			if pssble == n {
				if day7.Debug {
					fmt.Println(pssble)
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
