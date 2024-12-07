package day6

import (
	"bytes"
	"slices"

	"local/setutil"
)

type Pos struct {
	Row, Col int
}

var Rows, Cols int

var (
	Guard     Pos
	Obstacles = make(map[Pos]struct{})
)

func init() {
	nl := bytes.IndexByte(input, '\n')
	lines := slices.Collect(slices.Chunk(input, nl+1))

	Rows = len(lines)
	Cols = nl

	for row, line := range lines {
		line = line[:nl]
		for col, c := range line {
			switch c {
			case '#':
				setutil.Add(Obstacles, Pos{row, col})
			case '^':
				Guard = Pos{row, col}
			}
		}
	}
}
