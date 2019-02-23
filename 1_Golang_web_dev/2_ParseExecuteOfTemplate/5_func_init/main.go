package main

import (
	"log"
	"os"
	"text/template"
)

/*Create the template  */
var tpl *template.Template

// Init the template with pass a folder that contains all the files
func init() {
	/* Must is a wrapper function used to verifyt that a templateis valid
	 * during the parsing.*/
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	if err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil); err != nil {
		log.Fatalln(err)
	}

	if err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil); err != nil {
		log.Fatalln(err)
	}

	if err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil); err != nil {
		log.Fatalln(err)
	}
}
