package main

import (
	"fmt"

	"local/day1"
	"local/sliceutil"
)

func main() {
	l1, l2 := day1.Input()

	var similarity int

	for _, v := range l1 {
		similarity += v * sliceutil.Count(l2, v)
	}

	fmt.Println(similarity)
}
