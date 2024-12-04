package main

import (
	"fmt"
	"regexp"
	"strconv"

	"local/day3"
)

func main() {
	sum := 0

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	mtch := re.FindAllStringSubmatch(day3.Input, -1)

	for _, oprnds := range mtch {
		fmt.Println(oprnds)

		var err error
		var x, y int

		x, err = strconv.Atoi(oprnds[1])
		if err != nil {
			panic(err)
		}

		y, err = strconv.Atoi(oprnds[2])
		if err != nil {
			panic(err)
		}

		sum += x * y
	}

	fmt.Println(sum)
}
