package main

import (
	"html/template"
	"log"
	"os"
)

// Create the template
var tpl *template.Template

func init() {
	// Verify if the files is valid
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	// Passing the data to the template (index.gohtml) in the third param
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", 26)
	if err != nil {
		log.Fatalln(err)
	}
}
