package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// To create the default.html from main.go you need to writte in the command line: go run main > default.html
}
