package utils

import (
	"fmt"
	"math"
)

type inte interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect)area() float64 {
	return r.width * r.height
}

func (c circle) area() float64  {
	return math.Pi * c.radius * c.radius
}

func Inter()  {
	c := circle.area(circle{2})
	r := rect.area(rect{2,2})
	fmt.Println(c,r)
	c2 := circle{2}
	b := c2.area();

	fmt.Println(b)

}
