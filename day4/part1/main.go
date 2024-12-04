package main

import (
	"fmt"

	"local/day4"
)

func horizontal(input map[int]map[int]byte, row, col int) [4]byte {
	return [4]byte{input[row][col], input[row][col-1], input[row][col-2], input[row][col-3]}
}

func vertical(input map[int]map[int]byte, row, col int) [4]byte {
	return [4]byte{input[row][col], input[row-1][col], input[row-2][col], input[row-3][col]}
}

func forwardDiagonal(input map[int]map[int]byte, row, col int) [4]byte {
	return [4]byte{input[row][col], input[row-1][col+1], input[row-2][col+2], input[row-3][col+3]}
}

func backDiagonal(input map[int]map[int]byte, row, col int) [4]byte {
	return [4]byte{input[row][col], input[row-1][col-1], input[row-2][col-2], input[row-3][col-3]}
}

func word(s string) [4]byte {
	return [4]byte([]byte(s))
}

func main() {
	xmas, samx := word("XMAS"), word("SAMX")
	input, rows, cols := day4.Input, day4.Rows, day4.Cols

	found := make(map[int]map[int]struct{})
	for row := range rows {
		found[row] = make(map[int]struct{})
	}

	n := 0
	for row := range rows {
		for col := range cols {
			switch horizontal(input, row, col) {
			case xmas, samx:
				found[row][col] = struct{}{}
				found[row][col-1] = struct{}{}
				found[row][col-2] = struct{}{}
				found[row][col-3] = struct{}{}
				n++
			}
			switch vertical(input, row, col) {
			case xmas, samx:
				found[row][col] = struct{}{}
				found[row-1][col] = struct{}{}
				found[row-2][col] = struct{}{}
				found[row-3][col] = struct{}{}
				n++
			}
			switch forwardDiagonal(input, row, col) {
			case xmas, samx:
				found[row][col] = struct{}{}
				found[row-1][col+1] = struct{}{}
				found[row-2][col+2] = struct{}{}
				found[row-3][col+3] = struct{}{}
				n++
			}
			switch backDiagonal(input, row, col) {
			case xmas, samx:
				found[row][col] = struct{}{}
				found[row-1][col-1] = struct{}{}
				found[row-2][col-2] = struct{}{}
				found[row-3][col-3] = struct{}{}
				n++
			}
		}
	}

	for row := range rows {
		for col := range cols {
			c := '.'
			if _, found := found[row][col]; found {
				c = rune(input[row][col])
			}
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}

	fmt.Println(n)
}
