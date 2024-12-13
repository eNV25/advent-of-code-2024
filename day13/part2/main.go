package main

import (
	"fmt"
	"local/day13"

	"github.com/alex-ant/gomath/gaussian-elimination"
	"github.com/alex-ant/gomath/rational"
)

func main() {
	var price int64

	ps := day13.Input()

	for _, p := range ps {
		res, err := gaussian.SolveGaussian([][]rational.Rational{
			{rational.New(int64(p.ButtonA.X), 1), rational.New(int64(p.ButtonB.X), 1), rational.New(int64(p.Prize.X)+10000000000000, 1)},
			{rational.New(int64(p.ButtonA.Y), 1), rational.New(int64(p.ButtonB.Y), 1), rational.New(int64(p.Prize.Y)+10000000000000, 1)},
		}, day13.Debug)
		if err != nil {
			panic(err)
		}
		if day13.Debug {
			fmt.Println("->", res[0][0], res[1][0])
		}
		a, b := res[0][0], res[1][0]
		a.Simplify()
		b.Simplify()
		if a.IsNatural() && b.IsNatural() {
			price += a.GetNumerator()*3 + b.GetNumerator()
		}
	}

	fmt.Println(price)
}
