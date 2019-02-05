package main

import (
	"fmt"
	"math"
)

/*Rectangule ...-----------------------------------*/
type Rectangule struct {
	width  float32
	height float32
}

func (r Rectangule) area() float32 {
	return r.width * r.height
}

func (r Rectangule) perimeter() float32 {
	return (2 * r.width) + (2 * r.height)
}

/*Circle ...----------------------------------------*/
type Circle struct {
	radius float32
}

func (c Circle) area() float32 {
	return math.Pi * (c.radius * c.radius)
}

func (c Circle) perimeter() float32 {
	return (2 * math.Pi) * c.radius
}

/*geometry interface ---------------------------------*/
type geometry interface {
	area() float32
	perimeter() float32
}

/**
 * Main Function --------------------------------------*/
func main() {
	r := Rectangule{
		width:  3,
		height: 4,
	}

	c := Circle{radius: 5}

	measure(r)
	measure(c)
}

/*Function to provide the functionality of the interfaces*/
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("Area: ", g.area())
	fmt.Println("Perimeter: ", g.perimeter())
}
