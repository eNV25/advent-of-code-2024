package day9

import (
	"bytes"
	"slices"
)

type File struct {
	ID, Size int
}

func atoi(b byte) int {
	// itoa n = '0' + n
	return int(b - '0')
}

func Input() (files []File) {
	i := 0
	for pair := range slices.Chunk(bytes.TrimSpace(input), 2) {
		switch len(pair) {
		case 0:
		case 1:
			files = append(files, File{+i, atoi(pair[0])})
		case 2:
			files = append(files, File{+i, atoi(pair[0])})
			files = append(files, File{-1, atoi(pair[1])})
		}
		i++
	}
	return
}
