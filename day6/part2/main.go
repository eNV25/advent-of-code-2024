package main

import (
	"fmt"
	"iter"
	"maps"

	"local/day6"
	"local/setutil"
)

type Dir int

const (
	North Dir = iota
	East
	South
	West

	firstDir = North
	lastDir  = West
)

type Pos struct {
	day6.Pos
	Dir Dir
}

func (pos Pos) TurnRight() Pos {
	switch pos.Dir {
	case lastDir:
		pos.Dir = firstDir
	default:
		pos.Dir++
	}
	return pos
}

func (pos Pos) Next() Pos {
	switch pos.Dir {
	case North:
		pos.Row--
	case East:
		pos.Col++
	case South:
		pos.Row++
	case West:
		pos.Col--
	}
	return pos
}

func (pos Pos) IsInBounds() bool {
	return (0 <= pos.Row && pos.Row < day6.Rows) && (0 <= pos.Col && pos.Col < day6.Cols)
}

func Walk(pos Pos, obstacles map[day6.Pos]struct{}) iter.Seq2[Pos, bool] {
	return func(yield func(pos Pos, loops bool) bool) {
		turns := make(map[Pos]struct{})
		for {
			// When we detect a loop, return early.
			if setutil.Contains(turns, pos) {
				yield(pos, true)
				return
			}
			// When there is an obstacle next, turn right.
			// Else when we are out of bounds, return.
			// Else, go to the next position.
			if npos := pos.Next(); setutil.Contains(obstacles, npos.Pos) {
				setutil.Add(turns, pos)
				pos = pos.TurnRight()
			} else if !npos.IsInBounds() {
				return
			} else {
				pos = npos
			}
			// Yield all changes in position, including turns.
			if !yield(pos, false) {
				return
			}
		}
	}
}

func main() {
	loopObstacles := make(map[day6.Pos]struct{})

	for obs := range Walk(Pos{day6.Guard, North}, day6.Obstacles) {
		var npos Pos
		var loops bool
		var visited map[Pos]struct{}

		if obs.Pos == day6.Guard {
			continue
		}

		if _, found := loopObstacles[obs.Pos]; found {
			continue
		}

		if day6.Debug {
			visited = make(map[Pos]struct{})
		}

		obstacles := maps.Clone(day6.Obstacles)
		setutil.Add(obstacles, obs.Pos)
		for npos, loops = range Walk(Pos{day6.Guard, North}, obstacles) {
			if day6.Debug {
				setutil.Add(visited, npos)
			}
		}

		if loops {
			if day6.Debug {
				printMap(obs, visited)
			}
			setutil.Add(loopObstacles, obs.Pos)
		}
	}

	fmt.Println(len(loopObstacles))
}

func printMap(obs Pos, visited map[Pos]struct{}) {
	line := make([]byte, 0, day6.Cols)
	var pos day6.Pos
	for pos.Row = range day6.Rows {
		for pos.Col = range day6.Cols {
			c := byte('.')
			switch vert, hori := //
				setutil.Contains(visited, Pos{pos, North}) ||
					setutil.Contains(visited, Pos{pos, South}),
				setutil.Contains(visited, Pos{pos, East}) ||
					setutil.Contains(visited, Pos{pos, West}); {
			case pos == day6.Guard:
				c = '^'
			case pos == obs.Pos:
				c = 'O'
			case setutil.Contains(day6.Obstacles, pos):
				c = '#'
			case vert && hori:
				c = '+'
			case vert:
				c = '|'
			case hori:
				c = '-'
			}
			line = append(line, c)
		}
		fmt.Printf("%s\n", line)
		line = line[:0]
	}
	fmt.Println()
}
