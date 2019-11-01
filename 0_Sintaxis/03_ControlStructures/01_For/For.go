package main

import "fmt"

func main() {
	/* Ways to implement for statment */

	// first way
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	// Second way
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

}
