package main

import (
	"fmt"
	"math"
)

//Square ----------------------------------------------------
type square struct {
	side float32
}

func (s square) getArea() float32 {
	return s.side * s.side
}

//Circle ----------------------------------------------------
type circle struct {
	radius float32
}

func (c circle) getArea() float32 {
	return math.Pi * (c.radius * c.radius)
}

//Interface -------------------------------------------------
type shape interface {
	getArea() float32
}

//Interface GET INFO ----------------------------------------
func info(s shape) {
	fmt.Println(s.getArea())
}

/* Main function*/
func main() {
	s := square{side: 3}
	c := circle{radius: 5}

	info(s)
	info(c)
}
