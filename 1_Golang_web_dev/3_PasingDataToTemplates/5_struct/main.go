package main

import (
	"html/template"
	"log"
	"os"
)

type player struct {
	Name string
	No   int
	Team string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	players := []player{
		{
			Name: "Ronaldinho",
			No:   10,
			Team: "Brazil",
		}, {
			Name: "ronaldo",
			No:   10,
			Team: "Brazil",
		},
	}

	if err := tpl.ExecuteTemplate(os.Stdout, "simple.gohtml", players[0]); err != nil {
		log.Fatal(err)
	}

	if err := tpl.ExecuteTemplate(os.Stdout, "complex.gohtml", players); err != nil {
		log.Fatal(err)
	}
}
