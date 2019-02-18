package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name string
	age  int
}

type ByName []Person

func (this ByName) Len() int {
	return len(this)
}

func (this ByName) Less(i, j int) bool {
	return this[i].name < this[j].name
}

func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main() {
	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
	}

	sort.Sort(ByName(kids))
	fmt.Println(ByName(kids))
}
