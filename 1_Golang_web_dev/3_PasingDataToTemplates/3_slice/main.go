package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	players := []string{"Ronaldo", "Beckam", "Zidane", "Ronaldinho"}

	err := tpl.ExecuteTemplate(os.Stdout, "simple.gohtml", players)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "complex.gohtml", players)
	if err != nil {
		log.Fatalln(err)
	}
}
