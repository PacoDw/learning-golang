package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	/*Template is a container that contains all the html files (whatever it is) inside itself.
	 * ParseFiles: get n names of files to store them returning the *template (container) */
	tlp, err := template.ParseFiles("one.gmao")
	if err != nil {
		log.Fatal(err)
	}
	/*Execute parse ONLY the first file of the *template and passes the content in a writer as the first parameter
	 *the parameter could be a fille for example */
	err = tlp.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}

	tlp, err = tlp.ParseFiles("two.gmao", "vespa.gmao")
	if err != nil {
		log.Fatal(err)
	}

	/*It executes the selected file that is on the second parameter, it is executed if the template contains it.*/
	err = tlp.ExecuteTemplate(os.Stdout, "vespa.gmao", nil)
	if err != nil {
		log.Fatal(err)
	}

	if err = tlp.ExecuteTemplate(os.Stdout, "two.gmao", nil); err != nil {
		log.Fatal(err)
	}

	if err = tlp.ExecuteTemplate(os.Stdout, "one.gmao", nil); err != nil {
		log.Fatal(err)
	}

	/*Ever executes the first element inside of the template  */
	if err = tlp.Execute(os.Stdout, nil); err != nil {
		log.Fatal(err)
	}
}
