package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeDetails(s Shape) {
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}

func main() {
	c := Circle{radius: 5.0}
	printShapeDetails(c)
}
