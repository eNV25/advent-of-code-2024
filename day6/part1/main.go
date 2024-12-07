package main

import (
	"fmt"

	"local/day6"
	"local/setutil"
)

func main() {
	nvisited := 0

	pos := day6.Guard
	dir := up
	visited := make(map[day6.Pos]struct{})

	for (0 <= pos.Row && pos.Row < day6.Rows) &&
		(0 <= pos.Col && pos.Col < day6.Cols) {
		if !setutil.Contains(visited, pos) {
			nvisited++
		}

		setutil.Add(visited, pos)

		if setutil.Contains(day6.Obstacles, next(pos, dir)) {
			switch dir {
			case last:
				dir = first
			default:
				dir++
			}
		}

		pos = next(pos, dir)
	}

	fmt.Println(nvisited)
}

const (
	up = iota
	right
	down
	left

	first = up
	last  = left
)

func next(guard day6.Pos, dir int) day6.Pos {
	switch dir {
	case up:
		guard.Row -= 1
	case down:
		guard.Row += 1
	case left:
		guard.Col -= 1
	case right:
		guard.Col += 1
	}
	return guard
}
