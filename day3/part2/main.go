package main

import (
	"fmt"
	"regexp"
	"strconv"

	"local/day3"
)

func main() {
	sum := 0

	re := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\)`)

	mtch := re.FindAllStringSubmatch(day3.Input, -1)

	do := true
	for _, op := range mtch {
		fmt.Println(do, op)

		switch op[0] {
		case "do()":
			do = true
			continue
		case "don't()":
			do = false
			continue
		default: // mul(x,y)
			if !do {
				continue
			}
		}

		var err error
		var x, y int

		x, err = strconv.Atoi(op[1])
		if err != nil {
			continue
		}

		y, err = strconv.Atoi(op[2])
		if err != nil {
			continue
		}

		sum += x * y
	}

	fmt.Println(sum)
}
