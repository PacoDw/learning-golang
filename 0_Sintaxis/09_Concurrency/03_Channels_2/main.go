package main

import (
	"fmt"
	"time"
)

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

// this sents to the channel c chan<-
func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

// this receives from the channel c <-chan
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	// var c chan string = make(chan string)
	c := make(chan string)

	// We pass the channet to comunicate whatever message
	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
