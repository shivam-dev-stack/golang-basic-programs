package main

import (
	"fmt"
	"math"
)

// Shape interfacetype
type Shape interface {
	Area() float64
	Name() string
}

type Circle struct {
	Radius float64
}
type Rectangle struct {
	Length, Width float64
}
type Square struct {
	Side float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}
func (s Square) Area() float64 {
	return s.Side * s.Side
}

func (c Circle) Name() string    { return "Circle" }
func (r Rectangle) Name() string { return "Rectangle" }
func (s Square) Name() string    { return "Square" }

func PrintArea(s Shape) {
	fmt.Printf("The area of the %s is %.2f\n", s.Name(), s.Area())
}

func main() {

	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Length: 4, Width: 6},
		Square{Side: 3},
	}
	for _, shape := range shapes {
		PrintArea(shape)
	}
}
