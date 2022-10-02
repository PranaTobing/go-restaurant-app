package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float64
	getPerimeter() float64
}

type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) getArea() float64 {
	return r.width * r.height
}

func (r rectangle) getPerimeter() float64 {
	return r.width*2 + r.height*2
}

type square struct {
	width float64
}

func (s square) getArea() float64 {
	return s.width * s.width
}
func (r square) getPerimeter() float64 {
	return r.width * 4
}

type circle struct {
	radius float64
}

func (c circle) getArea() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) getPerimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measureShape(s shape) {
	fmt.Printf("Shape has area of %f\n", s.getArea())
	fmt.Printf("Shape has perimeter of %f\n", s.getPerimeter())
}

func main() {
	s := square{width: 5}
	r := rectangle{width: 10, height: 3}
	c := circle{radius: 7}

	measureShape(s)
	measureShape(r)
	measureShape(c)
}
