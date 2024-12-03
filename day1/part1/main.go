package main

import (
	"fmt"
	"slices"

	"local/day1"
	"local/iterutil"
)

func main() {
	s1, s2 := day1.Input()

	slices.Sort(s1)
	slices.Sort(s2)

	sum := 0

	for v1, v2 := range iterutil.Zip(s1, s2) {
		sum += max(v1-v2, v2-v1)
	}

	fmt.Println(sum)
}
