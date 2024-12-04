package day4

import (
	"bytes"
	_ "embed"
	"maps"
	"slices"
)

//go:embed example_input
var example_input []byte

//go:embed input
var input []byte

var (
	Input      map[int]map[int]byte
	Rows, Cols int
)

func init() {
	// input = example_input

	nl := bytes.IndexByte(input, '\n')

	Rows = nl
	Cols = nl

	Input = make(map[int]map[int]byte)
	Input[-3] = make(map[int]byte)
	Input[-2] = make(map[int]byte)
	Input[-1] = make(map[int]byte)
	Input[nl-1+1] = make(map[int]byte)
	Input[nl-1+2] = make(map[int]byte)
	Input[nl-1+3] = make(map[int]byte)

	for row, line := range slices.Collect(slices.Chunk(input, nl+1)) {
		Input[row] = make(map[int]byte)
		Input[row] = maps.Collect(slices.All(line[:nl]))
	}
}
