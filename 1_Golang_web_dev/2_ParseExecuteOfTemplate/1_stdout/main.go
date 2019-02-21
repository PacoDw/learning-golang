package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	/*Template is a container that contains all the html files (whatever it is) inside itself.
	 * ParseFiles: get n names of files to store them returning the *template (container)
	 */
	tpl, err := template.ParseFiles("1-file.gohtml", "1-file.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	/*Execute parse all the templates and passes the content in a writer as the firs parameter
	 *the parameter could be a fille for example
	 */
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}
}
