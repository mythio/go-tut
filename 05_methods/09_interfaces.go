package main

import (
	"fmt"
)

type geometry interface {
	area() float32
}

type Square struct {
	side float32
}

type Triangle struct {
	base   float32
	height float32
}

type Circle struct {
	radius float32
}

func (s Square) area() float32 {
	return s.side * s.side
}

func (c Circle) area() float32 {
	return 22.0 / 7.0 * c.radius
}

func (t Triangle) area() float32 {
	return 0.5 * t.base * t.height
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
}

func main() {
	s := Square{side: 2}
	t := Triangle{base: 3, height: 4}
	c := Circle{radius: 5}

	measure(s)
	measure(t)
	measure(c)
}
