package day5

import (
	"bufio"
	"embed"
	"os"
	"strconv"
	"strings"
)

//go:embed *input
var inputFS embed.FS

var (
	rules   = make(map[[2]int]struct{})
	Updates [][]int
)

func Cmp(a, b int) int {
	if a == b {
		return +0
	} else if _, less := rules[[2]int{a, b}]; less {
		return -1
	} else {
		return +1
	}
}

func init() {
	input, err := inputFS.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(input)

	for s.Scan() {
		line := s.Text()

		if line == "" {
			break
		}

		before, after, found := strings.Cut(line, "|")
		if !found {
			panic(found)
		}

		x, err := strconv.Atoi(before)
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(after)
		if err != nil {
			panic(err)
		}

		rules[[2]int{x, y}] = struct{}{}
	}

	for s.Scan() {
		line := strings.Split(s.Text(), ",")

		update := make([]int, 0, len(line))
		for _, v := range line {
			n, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			update = append(update, n)
		}

		Updates = append(Updates, update)
	}
}
