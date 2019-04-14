package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		if err := tpl.ExecuteTemplate(res, "index.gohtml", nil); err != nil {
			log.Fatalln()
		}
	})

	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))

	http.ListenAndServe(":3000", nil)
}
