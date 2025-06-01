package assorted

import (
	"fmt"
	"math"
)

// Rectange struct
type Rectange struct {
	length, width float64
}

// Circle struct
type Circle struct {
	radius float64
}

func (r Rectange) area() float64 {
	return r.length * r.width
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// AreaMain will be invoked from main.go
func AreaMain() {
	r1 := Rectange{12, 2}
	r2 := Rectange{9, 4}
	c1 := Circle{10}
	c2 := Circle{10}

	fmt.Println("Area of r1 =", r1.area())
	fmt.Println("Area of r2 =", r2.area())
	fmt.Println("Area of c1 =", c1.area())
	fmt.Println("Area of c2 =", c2.area())
}
