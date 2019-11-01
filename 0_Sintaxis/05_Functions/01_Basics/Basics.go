package main

import "fmt"

func main() {
	// res := []float32{1, 2, 3, 4, 5}

	// You can do this that
	fmt.Println(average([]float32{1, 2, 3, 4, 5}))

	// Other function called f2()
	fmt.Println(f2())
}

func average(xs []float32) float32 {
	var total float32

	for _, v := range xs {
		total += v
	}
	return total / float32(len(xs))
}

// We can also the retutn name
func f2() (name string) {
	name = "paco"
	return
}
