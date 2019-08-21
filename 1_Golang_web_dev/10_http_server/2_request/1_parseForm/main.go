package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type burger int

func (b burger) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}
func main() {
	var b burger

	http.ListenAndServe(":3000", b)
}
