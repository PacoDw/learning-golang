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

	type user struct {
		Name    string
		Motto   string
		IsAdmin bool
	}

	users := []user{
		{
			Name:    "Gandhi",
			Motto:   "Be the change",
			IsAdmin: true,
		},
		{
			Name:    "Buddha",
			Motto:   "The belief of no beliefs",
			IsAdmin: false,
		},
		{
			Name:    "",
			Motto:   "I dont know",
			IsAdmin: true,
		},
	}

	if err := tpl.Execute(os.Stdout, users); err != nil {
		log.Fatalln()
	}
}
