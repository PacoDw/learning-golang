package main

import (
	"fmt"
	"time"
)

func main() {
	// Creating two channels
	c1 := make(chan string)
	c2 := make(chan string)

	// Function sents to the channel c1 (chan<-)
	go func() {
		for {
			c1 <- "From 1"
		}
	}()

	// Function sents to the channel c2 (chan<-)
	go func() {
		for {
			c2 <- "From 2"
		}
	}()

	/*  Function receives from the channels (c1 & c2 ) the message
	and the Select statement select the channel that is alredy and show
	the message depends of the channel
	*/
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-time.After(time.Second):
				fmt.Println("timeout")
			default:
				fmt.Println("Nothing ready")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
