package main

import "fmt"

type Person struct {
	name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is ", p.name)
}

func main() {

	type Android struct {
		Person
		model string
	}

	a := Android{
		Person: Person{
			name: "Andy",
		},
		model: "k3455",
	}

	a.Talk()
}
