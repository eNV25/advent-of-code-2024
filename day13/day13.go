package day13

import (
	"bytes"
	"fmt"
	"io"
)

type Pos struct {
	X, Y int
}

func (r Pos) Add(o Pos) Pos {
	return Pos{X: r.X + o.X, Y: r.Y + o.Y}
}

func (r Pos) Subtract(o Pos) Pos {
	return Pos{X: r.X - o.X, Y: r.Y - o.Y}
}

type Prize struct {
	ButtonA, ButtonB, Prize Pos
}

func Input() (ps []Prize) {
	r := bytes.NewReader(input)

	for {
		var p Prize

		_, err := fmt.Fscanf(r, "Button A: X+%d, Y+%d\n", &p.ButtonA.X, &p.ButtonA.Y)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		}
		if err != nil {
			panic(err)
		}

		_, err = fmt.Fscanf(r, "Button B: X+%d, Y+%d\n", &p.ButtonB.X, &p.ButtonB.Y)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		}
		if err != nil {
			panic(err)
		}

		_, err = fmt.Fscanf(r, "Prize: X=%d, Y=%d\n", &p.Prize.X, &p.Prize.Y)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		}
		if err != nil {
			panic(err)
		}

		_, err = fmt.Fscanln(r)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		}
		if err != nil {
			panic(err)
		}

		ps = append(ps, p)
	}

	return
}
