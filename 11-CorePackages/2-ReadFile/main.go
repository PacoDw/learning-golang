package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	// Other way to get the current path
	// pwd, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// Getting the current path to get text.txt file
	dir, err := filepath.Abs("./text.txt")
	if err != nil {
		return
	}

	// Read file with the longer way
	go readFile(dir)

	// Read file with the shorter way
	go readFileShorter(dir)

	input := ""
	fmt.Scanln(&input)
}

func readFileShorter(dir string) {
	bs, err := ioutil.ReadFile(dir)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}

func readFile(dir string) {
	file, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
		/* Handle the error here */
		return
	}
	defer file.Close()

	// Get the file size
	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		fmt.Println(err)
		return
	}

	str := string(bs)
	fmt.Println(str)
}
