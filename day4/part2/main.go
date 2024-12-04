package main

import (
	"fmt"

	"local/day4"
)

func forwardDiagonal(input map[int]map[int]byte, row, col int) [3]byte {
	return [3]byte{input[row-1][col+1], input[row][col], input[row+1][col-1]}
}

func backDiagonal(input map[int]map[int]byte, row, col int) [3]byte {
	return [3]byte{input[row-1][col-1], input[row][col], input[row+1][col+1]}
}

func word(s string) [3]byte {
	return [3]byte([]byte(s))
}

func main() {
	mas, sam := word("MAS"), word("SAM")
	input, rows, cols := day4.Input, day4.Rows, day4.Cols

	n := 0
	for row := range rows {
		for col := range cols {
			switch forwardDiagonal(input, row, col) {
			case mas, sam:
				switch backDiagonal(input, row, col) {
				case mas, sam:
					n++
				}
			}
		}
	}

	fmt.Println(n)
}
