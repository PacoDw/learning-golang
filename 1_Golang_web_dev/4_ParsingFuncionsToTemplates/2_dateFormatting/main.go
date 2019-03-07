package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

var fns = template.FuncMap{
	"formmatYYMMDD": formatData,
}

// Crreating the func that will format the date in the index.gohtml
func formatData(t time.Time) string {
	return t.Format("2006-02-01")
}

func init() {
	tpl = template.Must(template.New("").Funcs(fns).ParseFiles("index.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", time.Now())
	if err != nil {
		log.Fatalln()
	}
}
