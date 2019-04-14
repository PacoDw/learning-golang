package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	scores := struct {
		Score1 int
		Score2 int
	}{
		7,
		9,
	}

	if err := tpl.Execute(os.Stdout, scores); err != nil {
		log.Fatalln()
	}
}
