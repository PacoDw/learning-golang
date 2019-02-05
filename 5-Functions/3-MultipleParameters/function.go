package main

import "fmt"

func main() {
	// You could pass any parameters in two ways

	// 1
	fmt.Println(add(1, 2, 3, 4))

	// 2
	xs := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(add(xs...))
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
