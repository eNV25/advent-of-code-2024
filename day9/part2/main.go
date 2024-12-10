package main

import (
	"fmt"
	"local/day9"
	"slices"
)

func checksum(files []day9.File) (cksum int) {
	var total int
	for _, file := range files {
		if file.ID >= 0 {
			total += file.Size
		}
	}

	for i := len(files) - 1; i >= 0; i-- {
		if files[i].ID < 0 {
			continue
		}
		empty := slices.IndexFunc(files[:i], func(e day9.File) bool {
			return e.ID < 0 && e.Size >= files[i].Size
		})
		if empty < 0 {
			continue
		}
		diff := files[empty].Size - files[i].Size
		files[empty] = files[i]
		files[i].ID = -1
		if diff > 0 {
			files = slices.Insert(files, empty+1, day9.File{ID: -1, Size: diff})
		}
		// defragment free space by squishing it to previous adjacent free space
		for i := i; 1 <= i && i < len(files); i++ {
			for i < len(files) && files[i-1].ID < 0 && files[i].ID < 0 {
				files[i-1].Size += files[i].Size
				files = slices.Delete(files, i, i+1)
			}
		}
	}

	if day9.Debug {
		fmt.Println(files)
	}

	if day9.Debug {
		fmt.Println("total", total)
	}

	pos := 0
	for _, file := range files {
		for range file.Size {
			if file.ID > 0 {
				cksum += pos * file.ID
			}
			pos++
		}
	}
	return
}

func main() {
	fmt.Println(checksum(day9.Input()))
}
