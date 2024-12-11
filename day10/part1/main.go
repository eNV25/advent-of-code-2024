package main

import (
	"fmt"
	"iter"
	"local/day10"

	"github.com/ugurcsen/gods-generic/sets/hashset"
)

var (
	rows, cols int
	grid       [][]int
)

func walkAndFind(row, col, v int, visited *hashset.Set[[2]int]) iter.Seq[struct{}] {
	return func(yield func(struct{}) bool) {
		if visited == nil {
			visited = hashset.New[[2]int]()
		}

		visited.Add([2]int{row, col})

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
			if !visited.Contains([2]int{pos.row, pos.col}) {
				if 0 <= pos.row && pos.row < rows && 0 <= pos.col && pos.col < cols {
					if diff := grid[pos.row][pos.col] - grid[row][col]; diff == 1 {
						walkAndFind(pos.row, pos.col, v, visited)(yield)
					}
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
