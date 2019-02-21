package main

import "fmt"

func main() {
	num := 5
	zero(&num) // To pass the addres to reference a location in memory
	fmt.Println(num)
}

func zero(numPtr *int) {
	*numPtr = 0 // The point can access to the value of the reference that has been passed
}
