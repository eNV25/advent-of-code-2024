package day8

import (
	"bytes"
	"cmp"
	"slices"
)

type Pos struct {
	Row, Col int
}

func ComparePos(a, b Pos) int {
	if c := cmp.Compare(a.Row, b.Row); c != 0 {
		return c
	}
	if c := cmp.Compare(a.Col, b.Col); c != 0 {
		return c
	}
	return 0
}

var (
	Input      = make(map[Pos]byte)
	Rows, Cols int
)

func (pos Pos) InBounds() bool {
	return (0 <= pos.Row && pos.Row < Rows) && (0 <= pos.Col && pos.Col < Cols)
}

func init() {
	lf := bytes.IndexByte(input, '\n')
	grid := slices.Collect(slices.Chunk(input, lf+1))

	Rows = len(grid)
	Cols = lf // ignore "\n"

	var pos Pos
	for pos.Row = range Rows {
		for pos.Col = range Cols {
			Input[pos] = grid[pos.Row][pos.Col]
		}
	}
}
