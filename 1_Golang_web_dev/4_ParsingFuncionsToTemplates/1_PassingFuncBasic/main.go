package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

/* FuncMap receives the function names that will be used when the files are parsed,
You just need to implement the function name in the file that you want to use.  */
var fn = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": addYearString,
}

/* This is my custom function and then it will be use when the files are parsed */
func addYearString(s string) string {
	return s + " creating to Nintendo 64"
}

/* To implements the functions, it is necessary use Funcs and pass it FuncMap */
func init() {
	tpl = template.Must(template.New("").Funcs(fn).ParseFiles("index.gohtml"))
}

func main() {

	bestGames := []struct {
		Name string
		Year string
	}{
		{
			Name: "Zelda: ocarina of time",
			Year: "1997",
		},
		{
			Name: "Super Smash Bros",
			Year: "1999",
		},
		{
			Name: "Super Mario Bros",
			Year: "1996",
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", bestGames)
	if err != nil {
		log.Fatalln(err)
	}
}
