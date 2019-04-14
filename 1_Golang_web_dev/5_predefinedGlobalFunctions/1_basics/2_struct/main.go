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

	book := struct {
		Pages []string
		Name  string
	}{
		[]string{"zero", "one", "two", "three", "four"},
		"Games Of Thrones",
	}

	if err := tpl.Execute(os.Stdout, book); err != nil {
		log.Fatalln()
	}
}
