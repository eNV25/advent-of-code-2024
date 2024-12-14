package main

import (
	"fmt"
	"local/day14"
)

func mod(a, b int) int {
	return ((a % b) + b) % b
}

func main() {
	state := make(map[[2]int]int)

	for p, v := range day14.Input() {
		state[[2]int{mod(p[0]+v[0]*100, width), mod(p[1]+v[1]*100, height)}]++
	}

	nw, ne, sw, se := 0, 0, 0, 0
	var p [2]int
	for p[0] = range width {
		for p[1] = range height {
			switch {
			case p[0] < width/2 && p[1] < height/2:
				nw += state[p]
			case p[0] > width/2 && p[1] < height/2:
				ne += state[p]
			case p[0] < width/2 && p[1] > height/2:
				sw += state[p]
			case p[0] > width/2 && p[1] > height/2:
				se += state[p]
			}
		}
	}

	fmt.Println(nw * ne * sw * se)
}
