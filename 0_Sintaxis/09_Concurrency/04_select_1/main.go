package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1

	for {
		select {
		case c <- x: // It sends to the 'c' channel the 'x' value for each iteration
			x, y = y, x+y
		case <-quit: // When the 'quit' channel receives any value, closes the channels
			close(c)
			close(quit)
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	// create the chanels
	c := make(chan int)
	quit := make(chan int)

	// The gorotine starts but waits until the channels close
	go func() {
		for {
			select {
			case v, ok := <-c: // Validates if the 'c' channel stills open
				if ok {
					// show the received value
					fmt.Println(v)

					// if the 'c' channel gets 34 then sends 0 to the 'quit' channel to close the channels
					if v == 34 {
						quit <- 0
						return
					}
				}
			}
		}
	}()

	fibonacci(c, quit)
}
