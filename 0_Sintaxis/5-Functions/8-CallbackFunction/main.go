package main

import (
	"fmt"
	"log"
)

type foo func(string, string)

func bar(route string, callback foo) bool {

	callback("param", "param")
	return true
}

func main() {
	fmt.Println("Hello, playground")

	res := bar("Hello", func(arg1, arg2 string) {
		log.Println(arg1 + "_1")
		log.Println(arg2 + "_2")
	})

	fmt.Println("res: ", res)
}
