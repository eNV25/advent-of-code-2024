package main

import (
	"fmt"
	"iter"
	"local/setutil"
	"slices"

	. "local/day8"
)

func combinations(input []Pos) iter.Seq2[Pos, Pos] {
	return func(yield func(Pos, Pos) bool) {
		// remove duplicates i.e. cat input | sort -u
		slices.SortFunc(input, ComparePos)
		input = slices.Compact(input)
		for i := range input {
			for j := i + 1; j < len(input); j++ {
				if !yield(input[i], input[j]) {
					return
				}
			}
		}
	}
}

func main() {
	towers := make(map[byte][]Pos)
	antinodes := make(map[Pos]struct{})
	for pos, c := range Input {
		if c != '.' {
			towers[c] = append(towers[c], pos)
		}
	}
	for _, towers := range towers {
		for pos1, pos2 := range combinations(towers) {
			diff := Pos{
				Row: pos2.Row - pos1.Row,
				Col: pos2.Col - pos1.Col,
			}

			if Debug {
				fmt.Println(pos1, pos2, diff)
			}

			for {
				if Debug {
					fmt.Println("=>", pos1)
				}
				setutil.Add(antinodes, pos1)

				pos1.Row -= diff.Row
				pos1.Col -= diff.Col

				if !pos1.InBounds() {
					break
				}
			}

			for {
				if Debug {
					fmt.Println("=>", pos2)
				}
				setutil.Add(antinodes, pos2)

				pos2.Row += diff.Row
				pos2.Col += diff.Col

				if !pos2.InBounds() {
					break
				}
			}
		}
	}
	fmt.Println(len(antinodes))
}
