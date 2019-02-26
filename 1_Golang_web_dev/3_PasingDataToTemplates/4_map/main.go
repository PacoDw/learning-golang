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

	players := map[string]string{
		"Ronaldo":    "no. 9",
		"Beckam":     "no. 23",
		"Zidane":     "no. 5",
		"Ronaldinho": "no. 10",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "simple.gohtml", players)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "complex.gohtml", players)
	if err != nil {
		log.Fatalln(err)
	}
}
