package main

import "fmt"

func main() {

	// fist example
	var i int
	increment := func() int {
		i++
		return i
	}
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println("-------------------------------")

	// -------------------------------------
	// Second example
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}

func makeEvenGenerator() func() int {
	i := 0
	return func() (ret int) {
		ret = i
		i += 2
		return
	}
}
