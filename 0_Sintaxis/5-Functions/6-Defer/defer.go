package main

import "fmt"

func main() {
	defer second()
	frist()
	third()
}

func frist() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func third() {
	fmt.Println("3rd")
}
