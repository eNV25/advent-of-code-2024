package day7

import (
	"bufio"
	"strconv"
	"strings"
)

var Input [][]int

func init() {
	s := bufio.NewScanner(strings.NewReader(input))

	for s.Scan() {
		var line []int

		before, after, found := strings.Cut(s.Text(), ": ")
		if !found {
			panic(`expected ": "`)
		}

		x, err := strconv.Atoi(before)
		if err != nil {
			panic(err)
		}
		line = append(line, x)

		for _, after := range strings.Split(after, " ") {
			y, err := strconv.Atoi(after)
			if err != nil {
				panic(err)
			}
			line = append(line, y)
		}

		Input = append(Input, line)
	}
}
