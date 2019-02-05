package main

import "fmt"

/* Create the Person Struct --------------- */
type Person struct {
	name string
}

func (p Person) getName() string {
	return p.name
}

/* End of the struct -----------------------*/

type MyInterface interface {
	getName() string
}

func main() {
	greet(Person{name: "Paco Preciado"})
}

func greet(i MyInterface) {
	fmt.Println(i.getName())
}
