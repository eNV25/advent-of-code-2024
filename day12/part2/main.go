package main

import (
	"fmt"
	"local/day12"
	set "local/setutil"
)

func visit(visited set.Set[[2]int], grid *day12.Grid, area, corners *int, row, col int) {
	visited.Add([2]int{row, col})
	*area++

	// NW
	if grid.X[row-1][col] == grid.X[row][col] && grid.X[row][col-1] == grid.X[row][col] && grid.X[row-1][col-1] != grid.X[row][col] {
		*corners++
	} else //
	if grid.X[row-1][col] != grid.X[row][col] && grid.X[row][col-1] != grid.X[row][col] && grid.X[row-1][col-1] != grid.X[row][col] {
		*corners++
	} else //
	if grid.X[row-1][col] != grid.X[row][col] && grid.X[row][col-1] != grid.X[row][col] {
		*corners++
	}
	// NE
	if grid.X[row-1][col] == grid.X[row][col] && grid.X[row][col+1] == grid.X[row][col] && grid.X[row-1][col+1] != grid.X[row][col] {
		*corners++
	} else //
	if grid.X[row-1][col] != grid.X[row][col] && grid.X[row][col+1] != grid.X[row][col] && grid.X[row-1][col+1] != grid.X[row][col] {
		*corners++
	} else //
	if grid.X[row-1][col] != grid.X[row][col] && grid.X[row][col+1] != grid.X[row][col] {
		*corners++
	}
	// SE
	if grid.X[row+1][col] == grid.X[row][col] && grid.X[row][col+1] == grid.X[row][col] && grid.X[row+1][col+1] != grid.X[row][col] {
		*corners++
	} else //
	if grid.X[row+1][col] != grid.X[row][col] && grid.X[row][col+1] != grid.X[row][col] && grid.X[row+1][col+1] != grid.X[row][col] {
		*corners++
	} else //
	if grid.X[row+1][col] != grid.X[row][col] && grid.X[row][col+1] != grid.X[row][col] {
		*corners++
	}
	// SW
	if grid.X[row+1][col] == grid.X[row][col] && grid.X[row][col-1] == grid.X[row][col] && grid.X[row+1][col-1] != grid.X[row][col] {
		*corners++
	} else //
	if grid.X[row+1][col] != grid.X[row][col] && grid.X[row][col-1] != grid.X[row][col] && grid.X[row+1][col-1] != grid.X[row][col] {
		*corners++
	} else //
	if grid.X[row+1][col] != grid.X[row][col] && grid.X[row][col-1] != grid.X[row][col] {
		*corners++
	}

	for _, neigh := range [...]struct{ row, col int }{
		{row - 1, col},
		{row, col + 1},
		{row + 1, col},
		{row, col - 1},
	} {
		if grid.X[neigh.row][neigh.col] == grid.X[row][col] {
			if !visited.Contains([2]int{neigh.row, neigh.col}) {
				visit(visited, grid, area, corners, neigh.row, neigh.col)
			}
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
				area, corners := 0, 0
				visit(visited, &grid, &area, &corners, row, col)
				if day12.Debug {
					fmt.Printf("%c -> area * corners = %d * %d = %d\n", grid.X[row][col], area, corners, area*corners)
				}
				price += area * corners
			}
		}
	}

	fmt.Println(price)
}
