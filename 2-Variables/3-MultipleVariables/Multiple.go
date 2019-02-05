package main

import "fmt"

func main() {
	// you could use const instead var
	var (
		greet = "Hello!"
		name  = "Paco"
	)

	fmt.Println(greet + " " + name)
}
