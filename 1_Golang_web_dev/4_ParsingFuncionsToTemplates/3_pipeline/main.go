package main

import (
	"html/template"
	"log"
	"math"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fns).ParseFiles("index.gohtml"))
}

func double(n int) int {
	return n + n
}

func square(n int) float64 {
	return math.Pow(float64(n), 2)
}

func sqRoot(n int) float64 {
	return math.Sqrt(float64(n))
}

var fns = template.FuncMap{
	"double": double,
	"square": square,
	"sqRoot": sqRoot,
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", 3)
	if err != nil {
		log.Fatalln()
	}
}
