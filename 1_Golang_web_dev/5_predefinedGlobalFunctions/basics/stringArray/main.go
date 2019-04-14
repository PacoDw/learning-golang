package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	stringArray := []string{"zero", "one", "two", "three", "four", "five"}

	if err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", stringArray); err != nil {
		log.Fatalln()
	}
}
