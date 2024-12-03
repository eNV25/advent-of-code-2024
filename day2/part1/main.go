package main

import (
	"fmt"

	"local/day2"
	"local/iterutil"
)

func main() {
	rs := day2.Input()
	safe := 0
	for _, r := range rs {
		inc := true
		dec := true
		for v1, v2 := range iterutil.Zip(r, r[1:]) {
			diff := max(v1-v2, v2-v1)
			if diff < 1 || diff > 3 {
				goto unsafe
			}
			if inc && v1 > v2 {
				if !dec {
					goto unsafe
				}
				inc = false
			}
			if dec && v1 < v2 {
				if !inc {
					goto unsafe
				}
				dec = false
			}
		}
		safe++
	unsafe:
	}
	fmt.Println(safe)
}
