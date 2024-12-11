package main

import (
	"fmt"
	"local/day11"
)

func main() {
	state := make(map[int]int)
	for _, e := range day11.Input() {
		state[e] = 1
	}

	if day11.Debug {
		fmt.Println(state)
	}

	for range func() int {
		if day11.Debug {
			return 25
		} else {
			return 75
		}
	}() {
		eval(state)
	}

	sum := 0
	for _, n := range state {
		sum += n
	}
	fmt.Println(sum)
}

type change struct{ k, v int }

var changes []change

func eval(state map[int]int) {
	changes = changes[:0]
	for e, v := range state {
		if day11.Debug {
			fmt.Println(e, v)
		}
		changes = append(changes, change{e, -v})
		if e == 0 {
			changes = append(changes, change{1, v})
		} else if a, b, ok := splitEvenDigits(e); ok {
			changes = append(changes, change{a, v})
			changes = append(changes, change{b, v})
		} else {
			changes = append(changes, change{e * 2024, v})
		}
	}
	for _, change := range changes {
		state[change.k] += change.v
	}
	for k, v := range state {
		if v <= 0 {
			delete(state, k)
		}
	}
	if day11.Debug {
		fmt.Println(changes)
	}
	if day11.Debug {
		fmt.Println(state)
	}
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
