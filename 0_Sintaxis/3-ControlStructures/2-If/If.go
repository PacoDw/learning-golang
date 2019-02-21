package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}

	// Other uses with if statement
	// if el, ok := elements["Li"]; ok {
	// 	fmt.Println(el["name"], el["state"])
	// }
}
