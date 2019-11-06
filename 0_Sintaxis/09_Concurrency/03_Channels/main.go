package main

import "fmt"

func main() {
	// we can add a buffer to limit the chanel
	c := make(chan int, 2)

	c <- 1
	c <- 2
	// c <- 3 // if you run with this code line it gets an error because the bueffer is full

	fmt.Println(<-c)
	fmt.Println(<-c)
	// fmt.Println(<-c) // with this line
}
