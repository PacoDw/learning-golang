package main

import "fmt"

func main() {
	// we can add a buffer to limit the chanel
	c := make(chan int, 2)

	c <- 1
	c <- 2
	go func() { c <- 3 }() // this doesn't get an error because the func is called before
	// that buffer is fulled and our code doesn't block the main thread

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
