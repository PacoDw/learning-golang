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
	err := tpl.Execute(os.Stdout, "This is my content")
	if err != nil {
		log.Fatalln(err)
	}
}
