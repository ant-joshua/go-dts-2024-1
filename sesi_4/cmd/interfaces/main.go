package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func main() {
	var r shape = rectangle{width: 3, height: 4}
	var c shape = circle{radius: 5}

	fmt.Printf("Type of r: %T\n", r)

	r.area()
	r.perimeter()

	fmt.Printf("Type of c: %T\n", c)

	c.area()
	c.perimeter()

	c.(circle).volume()
}
