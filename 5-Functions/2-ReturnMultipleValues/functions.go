package main

import "fmt"

func main() {

	// You need implement this example logic implementation
	if res, ok, err := validateNaturalNumber(1); ok {
		fmt.Println(res)
	} else {
		fmt.Println(err)
	}
}

func validateNaturalNumber(number byte) (num byte, ok bool, err string) {
	ok = false
	if number < 10 {
		ok = true
		num = number
	} else {
		err = "Err: The number is not a natural number"
	}
	return num, ok, err // You could returning multiple variables
}
