package main

import (
	"fmt"
	"slices"

	"local/day5"
)

func main() {
	sum := 0
	for _, update := range day5.Updates {
		if slices.IsSortedFunc(update, day5.Cmp) {
			sum += update[len(update)/2]
		}
	}
	fmt.Println(sum)
}
