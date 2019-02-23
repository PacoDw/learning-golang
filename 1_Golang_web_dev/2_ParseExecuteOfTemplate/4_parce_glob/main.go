package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	/*The template (container) with this function get itself all the files of the folder inside */
	tlp, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatalln(err)
	}

	//it executes the first file of the template (container)
	if err = tlp.Execute(os.Stdout, nil); err != nil {
		log.Fatalln(err)
	}

	//It executes the selected file
	if err = tlp.ExecuteTemplate(os.Stdout, "vespa.gmao", nil); err != nil {
		log.Fatalln(err)
	}

	if err = tlp.ExecuteTemplate(os.Stdout, "two.gmao", nil); err != nil {
		log.Fatalln(err)
	}

	if err = tlp.ExecuteTemplate(os.Stdout, "one.gmao", nil); err != nil {
		log.Fatalln(err)
	}
}
