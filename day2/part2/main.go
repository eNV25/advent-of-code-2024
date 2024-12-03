package main

import (
	"fmt"
	"slices"

	"local/day2"
	"local/iterutil"
	"local/sliceutil"
)

func isSafe(tolerate bool, report []int) bool {
	inc := true
	dec := true
	for v1, v2 := range iterutil.Zip(report, report[1:]) {
		if diff := max(v1-v2, v2-v1); diff < 1 || diff > 3 {
			if tolerate {
				for i := 0; i < len(report); i++ {
					if isSafe(false, sliceutil.Deleted(report, i, i+1)) {
						return true
					}
				}
			}
			return false
		}
		if inc && v1 > v2 {
			if !dec { // not init
				if tolerate {
					for i := 0; i < len(report); i++ {
						if isSafe(false, sliceutil.Deleted(report, i, i+1)) {
							return true
						}
					}
				}
				return false
			}
			inc = false
		}
		if dec && v1 < v2 {
			if !inc { // not init
				if tolerate {
					for i := 0; i < len(report); i++ {
						if isSafe(false, sliceutil.Deleted(report, i, i+1)) {
							return true
						}
					}
				}
				return false
			}
			dec = false
		}
	}
	return true
}

func main() {
	rs := day2.Input()
	safe := 0
	for r := range slices.Values(rs) {
		if isSafe(true, r) {
			safe++
		}
	}
	fmt.Println(safe)
}
