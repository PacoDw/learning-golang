package main

import "fmt"

func main() {
	// It create a new pointer
	xPtr := new(int)
	fmt.Println(*xPtr)
	one(xPtr)
	fmt.Println(*xPtr)

}

func one(nPtr *int) {
	*nPtr = 1
}
