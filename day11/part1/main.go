package main

import (
	"fmt"
	"local/day11"
)

func main() {
	state := day11.Input()

	for range 25 {
		state = eval(state)
	}

	fmt.Println(len(state))
}

func eval(state []int) []int {
	ret := state[:0:0]
	for _, e := range state {
		if e == 0 {
			ret = append(ret, 1)
		} else if a, b, ok := splitEvenDigits(e); ok {
			ret = append(ret, a, b)
		} else {
			ret = append(ret, e*2024)
		}
	}
	return ret
}

func splitEvenDigits(e int) (a int, b int, ok bool) {
	digits := 0
	for i := 1; i <= e; i *= 10 {
		digits++
	}
	if digits%2 != 0 {
		return 0, 0, false
	}
	b = 0
	d := 1
	for range digits / 2 {
		b, e = b+e%10*d, e/10
		d *= 10
	}
	a = e
	return a, b, true
}
