package main

import (
	"fmt"
	"strconv"
)

// Person Struct
type Person struct {
	name string
	age  byte
	job  string
}

// Function get full info of the Person struct
func (person *Person) getFullInfo() (info string) {
	info = "Name " + person.name + " with " + strconv.Itoa(int(person.age)) + " years old have a job as a " + person.job
	return
}

func main() {
	employee := Person{
		name: "Perenganito",
		age:  24,
		job:  "Programer",
	}

	fmt.Println(employee.getFullInfo())
}
