package day12

import (
	"bytes"
	"slices"
)

type Grid struct {
	X          map[int]map[int]byte
	Rows, Cols int
}

func (g *Grid) InBounds(row, col int) bool {
	return 0 <= row && row < g.Rows && 0 <= col && col < g.Cols
}

func Input() Grid {
	lf := bytes.IndexByte(input, '\n')
	sgrid := slices.Collect(slices.Chunk(input, lf+1))
	grid := make(map[int]map[int]byte)
	for i, row := range sgrid {
		grid[i] = make(map[int]byte)
		for j, b := range row[:lf] {
			grid[i][j] = b
		}
	}
	return Grid{Rows: len(sgrid), Cols: lf, X: grid}
}
