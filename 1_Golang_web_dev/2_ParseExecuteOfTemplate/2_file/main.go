package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	/*Template is a container that contains all the html files (whatever it is) inside itself.
	 * ParseFiles: get n names of files to store them returning the *template (container)
	 */
	tlp, err := template.ParseFiles("index.gohtml", "1-file.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	/*Create a new file called index.html
	 */
	file, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	/*Execute parse ONLY the first file of the *template and passes the content in a writer as the first parameter
	 *the parameter could be a fille for example
	 */
	err = tlp.Execute(file, nil)
	if err != nil {
		log.Fatal(err)
	}
}
