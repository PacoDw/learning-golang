package main

import (
	"fmt"
	"strconv"
)

/*Person ---------------------------------------- */
type person struct {
	name string
	age  int
}

func (p person) pSpeak() string {
	return `I am ` + p.name + ` and my years old is ` + strconv.Itoa(p.age)
}

/*Secret Agent ----------------------------------- */
type secretAgent struct {
	person
	secretName string
}

func (sa secretAgent) saSpeak() string {
	return `I am ` + sa.name + ` and my years old is ` + strconv.Itoa(sa.age) + ` my secret name is ` + sa.secretName
}

type implement interface {
	pSpeak() string
}

func speak(i implement) {
	fmt.Println(i.pSpeak())
}

/*MAIN FUNCTION */
func main() {
	p := person{
		name: "Paco",
		age:  26,
	}
	fmt.Println(p.name)

	sa := secretAgent{
		person: person{
			name: "James Bond",
			age:  27,
		},
		secretName: "007",
	}
	fmt.Println(sa.name)

	fmt.Println(p.pSpeak())

	fmt.Println(sa.pSpeak())
	fmt.Println(sa.saSpeak())

	// You can create a interface with get a structure like agent
	// that inherit of person
	speak(sa)
	speak(p)
}
