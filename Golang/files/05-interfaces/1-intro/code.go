package main

import (
	"fmt"
	"math"
)

/*
Interface
- like a contract
- need to fulfill some conditions
*/

type shape interface {
	area() float64
	perimeter() float64
}

type circle struct{
	radius float64
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func shapeInfo(s shape) (float64, float64) {
	return s.perimeter(), s.area()
}

func main() {
	ring := circle{radius: 10}
	a, b := shapeInfo(ring)
	fmt.Println(a, b)
}