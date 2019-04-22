package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				{ // example with key value
					Number: "CSCI-40",
					Name:   "Introduction to Programming in Go",
					Units:  "4",
				}, // examples without key
				{"CSCI-130", "Introduction to Web Programming with Go", "4"},
				{"CSCI-140", "Mobile Apps Using Go", "4"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				{"CSCI-50", "Advanced Go", "5"},
				{"CSCI-190", "Advanced Web Programming with Go", "5"},
				{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatalln(err)
	}
}
