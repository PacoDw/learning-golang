package main

import "fmt"

func main() {
	type Person struct {
		name string
		last string
		age  byte
		job  string
	}

	/* We can create an instance of our new Person type in variety of ways: */

	// var fulanito Person

	// fulatino := new(Person)

	fulanito := Person{
		name: "Fulanito Perenganito",
		last: "Del Toro",
		age:  22,
		job:  "Programer",
	}

	fmt.Println(fulanito)

}
