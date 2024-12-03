package day2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Input() (rows [][]int) {
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		if err := in.Err(); err != nil {
			panic(err)
		}

		fields := strings.Fields(in.Text())

		row := make([]int, 0, len(fields))

		for _, field := range fields {
			v, err := strconv.Atoi(field)
			if err != nil {
				panic(err)
			}
			row = append(row, v)
		}

		rows = append(rows, row)
	}
	return
}
