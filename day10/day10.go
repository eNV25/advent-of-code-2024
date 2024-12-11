package day10

import (
	"bytes"
	"slices"
)

func Input() (rows, cols int, grid [][]int) {
	lf := bytes.IndexByte(input, '\n')
	input := slices.Collect(slices.Chunk(input, lf+1))

	rows = len(input)
	cols = lf

	grid = make([][]int, rows)
	for row := range grid {
		grid[row] = make([]int, cols)
		for col := range grid[row] {
			grid[row][col] = atoi(input[row][col])
		}
	}

	return
}

func atoi(b byte) int {
	if '0' <= b && b <= '9' {
		return int(b - '0')
	}
	return 0
}
