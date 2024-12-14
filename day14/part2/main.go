package main

import (
	"fmt"
	"local/day14"
)

func mod(a, b int) int {
	return ((a % b) + b) % b
}

func state(n int) map[[2]int]int {
	state := make(map[[2]int]int)
	for p, v := range day14.Input() {
		state[[2]int{mod(p[0]+v[0]*n, width), mod(p[1]+v[1]*n, height)}]++
	}
	return state
}

func main() {
	for i := 1; i < 10000; i++ {
		fmt.Println("->", i)
		state := state(i)
		var p [2]int
		for p[0] = range width {
			for p[1] = range height {
				c := byte('.')
				if n := state[p]; 0 < n && n <= 9 {
					c = byte('#')
				}
				fmt.Printf("%c", c)
			}
			fmt.Println()
		}
	}
}
