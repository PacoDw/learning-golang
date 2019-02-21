package main

import "fmt"

func main() {
	/*Maps*/
	x := make(map[string]int)

	x["key"] = 10
	x["other"] = 3

	fmt.Println(x)

	// Delete function to remove any item
	delete(x, "other")
	fmt.Println(x)

	// Check if an element exists
	value, ok := x["other"] // other not
	fmt.Println(value, ok)

	// example to validate if exists
	if val, okay := x["key"]; okay {
		fmt.Println(val, okay)
	}

	// Shorther to create maps
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}
