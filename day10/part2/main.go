package main

import (
	"fmt"
	"iter"
	"local/day10"
)

var (
	rows, cols int
	grid       [][]int
)

func walkAndFind(row, col, v int, path [][2]int) iter.Seq[struct{}] {
	return func(yield func(struct{}) bool) {
		path = append(path, [2]int{row, col})

		if day10.Debug {
			fmt.Println(path, grid[row][col])
		}

		if grid[row][col] == v {
			yield(struct{}{})
			return
		}

		for _, pos := range [...]struct{ row, col int }{
			{row, col + 1},
			{row + 1, col},
			{row, col - 1},
			{row - 1, col},
		} {
			if 0 <= pos.row && pos.row < rows && 0 <= pos.col && pos.col < cols {
				if diff := grid[pos.row][pos.col] - grid[row][col]; diff == 1 {
					n := len(path)
					walkAndFind(pos.row, pos.col, v, path)(yield)
					path = path[:n]
				}
			}
		}
	}
}

func main() {
	rows, cols, grid = day10.Input()

	for row := range grid {
		if day10.Debug {
			fmt.Println(grid[row])
		}
	}

	sum := 0

	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 0 {
				for range walkAndFind(row, col, 9, nil) {
					sum += 1
				}
			}
		}
	}

	fmt.Println(sum)
}
