package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// This will be embbebed into the html content
	name := "Paco Preciado Mendoza"

	// The content that the file will contain
	str := fmt.Sprint(`
					<!DOCTYPE html>
					<html lang="en">
					<head>
						<meta charset="UTF-8">
						<title>Hello World!</title>
					</head>
					<body>
						<h1>` + name + `</h1>
					</body>
					</html>
	`)

	// Creating the file
	file, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Error creating file", err)
	}
	defer file.Close()

	// Copy the content into the index.html file
	io.Copy(file, strings.NewReader(str))
}
