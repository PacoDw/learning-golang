package main

import "os"

func main() {
	file, err := os.Create("Text_test.txt")
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString("Test content")
}
