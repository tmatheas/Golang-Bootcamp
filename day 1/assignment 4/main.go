package main

import "math"

type Shape interface {
	getArea() float64
}

type Triange struct {
	base, height float64
}

type Square struct {
	side float64
}

func (t Triange) getArea() float64 {
	return t.base * t.height * 0.5
}

func (s Square) getArea() float64 {
	return math.Pow(s.side, 2)
}

func printArea(s Shape) {
	println("Area: ", s.getArea(), " sq.units")
}

func main() {
	triangle := Triange{base: 5, height: 6}
	square := Square{side: 7}
	printArea(triangle)
	printArea(square)
}
