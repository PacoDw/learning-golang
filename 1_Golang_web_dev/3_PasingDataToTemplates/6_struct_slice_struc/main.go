package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type player struct {
	Name string
	No   int
}

type team struct {
	Name      string
	Years     int
	NoPlayers int
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	players := []player{
		{
			Name: "ronaldo",
			No:   9,
		},
		{
			Name: "ronaldinho",
			No:   10,
		},
		{
			Name: "beckam",
			No:   23,
		},
	}

	teams := []team{
		{
			Name:      "Barcelona",
			Years:     130,
			NoPlayers: 45,
		},
		{
			Name:      "Real Madrid",
			Years:     140,
			NoPlayers: 49,
		},
		{
			Name:      "Manchester City",
			Years:     137,
			NoPlayers: 32,
		},
	}

	items := struct {
		Jugadores []player
		Equipos   []team
	}{
		Jugadores: players,
		Equipos:   teams,
	}

	if err := tpl.ExecuteTemplate(os.Stdout, "simple.gohtml", items); err != nil {
		log.Fatalln(err)
	}

	if err := tpl.ExecuteTemplate(os.Stdout, "complex.gohtml", items); err != nil {
		log.Fatalln(err)
	}
}
