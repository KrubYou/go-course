package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}
type react struct {
	width, height float64
}
type circle struct {
	redius float64
}

func (r react) area() float64 {
	return r.width * r.height
}

func (r react) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.redius * c.redius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.redius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("Area =", g.area())
	fmt.Println("perimeter =", g.perim())
}

func main() {

	r := react{width: 10, height: 20}
	c := circle{redius: 10}
	measure(r)
	measure(c)
}
