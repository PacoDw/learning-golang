package main

import "fmt"

func main() {
	/* Ways to use arrays*/

	// ----------------------------------
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	// ----------------------------------
	y := [5]float32{98, 92, 77, 65, 85}

	var total float32
	for i := 0; i < len(y); i++ {
		total += y[i]
	}
	fmt.Println(total, "/", len(y), "=", total/float32(len(y)))

	// ----------------------------------
	z := [5]float32{34, 56, 67, 84, 69}
	var res float32

	for _, value := range z {
		res += value
	}
	fmt.Println(res, "/", float32(len(z)), "=", res/float32(len(z)))
}
