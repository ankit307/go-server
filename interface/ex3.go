package main

import (
	"fmt"
	"math"
)

type geometery interface {
	area() float64
	perim() float64
}

type rect struct {
	height, width float64
}

type circle struct {
	radius float64
}

func (r *rect) area() float64 {
	return r.height * r.width
}

func (r *rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	rect := rect{height: 20, width: 10}
	fmt.Printf("Area of react: %f\n", rect.area())
	fmt.Printf("Perimeter is :%f\n", rect.perim())

	circle := circle{radius: 20}
	fmt.Printf("Area of circle : %f\n", circle.area())
	fmt.Printf("Perimeter is: %f\n", circle.perim())
}
