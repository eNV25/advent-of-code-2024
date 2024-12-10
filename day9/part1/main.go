package main

import (
	"fmt"
	"local/day9"
)

func checksum(files []day9.File) (cksum int) {
	var total int
	for _, file := range files {
		if file.ID >= 0 {
			total += file.Size
		}
	}
	start := 0
	end := len(files) - 1
	if day9.Debug {
		fmt.Println("total", total)
	}
	for pos := range total {
		for start < len(files)-1 && files[start].Size <= 0 {
			start++
		}
		files[start].Size--

		if empty := files[start].ID < 0; !empty {
			if day9.Debug {
				fmt.Println(pos, "cksum start", files[start])
			}
			cksum += pos * files[start].ID
		} else {
			for end >= 1 && (files[end].ID < 0 || files[end].Size <= 0) {
				end--
			}
			files[end].Size--

			if day9.Debug {
				fmt.Println(pos, "cksum end", files[end])
			}
			cksum += pos * files[end].ID
		}
	}
	return
}

func main() {
	fmt.Println(checksum(day9.Input()))
}
