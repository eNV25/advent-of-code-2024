package main

import (
	"fmt"
	"local/day12"
	set "local/setutil"
)

func visit(visited set.Set[[2]int], grid *day12.Grid, area, perimeter *int, row, col int) {
	visited.Add([2]int{row, col})
	*area++

	for _, neigh := range [...]struct{ row, col int }{
		{row - 1, col},
		{row, col + 1},
		{row + 1, col},
		{row, col - 1},
	} {
		if grid.X[neigh.row][neigh.col] == grid.X[row][col] {
			if !visited.Contains([2]int{neigh.row, neigh.col}) {
				visit(visited, grid, area, perimeter, neigh.row, neigh.col)
			}
		} else {
			*perimeter++
		}
	}
}

func main() {
	grid := day12.Input()
	visited := make(set.Set[[2]int])

	price := 0

	for row := range grid.X {
		for col := range grid.X[row] {
			if !visited.Contains([2]int{row, col}) {
				area, perimeter := 0, 0
				visit(visited, &grid, &area, &perimeter, row, col)
				if day12.Debug {
					fmt.Printf("%c -> area * perimeter = %d * %d = %d\n", grid.X[row][col], area, perimeter, area*perimeter)
				}
				price += area * perimeter
			}
		}
	}

	fmt.Println(price)
}
