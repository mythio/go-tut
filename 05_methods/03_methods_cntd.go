package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

type MyFloat float64

func (p Point) Print() {
	fmt.Println(p)
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	p := Point{1, 2}
	p.Print()
}
