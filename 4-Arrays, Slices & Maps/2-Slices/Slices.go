package main

import "fmt"

func main() {
	/*To create a slice you need to use the built-in make function*/
	x := make([]float32, 5) // allows 3 parameters
	fmt.Println(x)

	// other way to create
	arr := [5]float32{1, 2, 3, 4, 5}
	y := arr[0:2]
	fmt.Println(y)

	/*Slice Functions*/
	slice1 := []int{1, 2, 3}

	// append function
	slice2 := append(slice1, 4, 5, 7)
	fmt.Println(slice1, "append to", slice2)

	// copy function
	slice3 := make([]int, 5)
	copy(slice3, slice1)
	fmt.Println(slice1, "copy to", slice3)

}
