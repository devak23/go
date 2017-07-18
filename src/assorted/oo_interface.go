package assorted

import (
	"fmt"
	"math"
)

// define an interface
type Shape interface {
	area() float64
	name() string
}

// define a circle object
type GCircle struct {
	radius float64
}

// define a square object
type GSquare struct {
	side float64
}

// implement the Shape interface's method name() for circle
func (c *GCircle) name() string {
	return "MyCircle"
}

// implement the Shape interface's method name() for Square
func (s *GSquare) name() string {
	return "MySquare"
}

// implement the Shape interface's method area() for circle
func (c *GCircle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// implement the Shape interface's method area() for Square
func (s *GSquare) area() float64 {
	return math.Pow(s.side, 2)
}

// The main method
func OOInterfaceMain() {
	myCircle := GCircle{22.0}
	mySquare := GSquare{22.0}

	calculateArea(&myCircle, &mySquare)
}

// Helper method which takes a varargs of Shape objects
// and prints their area
func calculateArea(shapes ...Shape) {
	for _, v := range shapes {
		fmt.Printf("Area of %s is %f\n", v.name(), v.area())
	}
}
