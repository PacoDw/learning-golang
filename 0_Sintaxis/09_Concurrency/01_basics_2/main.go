package main

import (
	"fmt"
	"time"
)

func say(greet string) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("index: %d message: %s \n", i, greet)
	}
}

func main() {
	go say("Hola")
	say("bye")
}
