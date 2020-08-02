package main

import (
	"os"
	"log"
	"text/template"
)

var tpl *template.Template

type Game struct {
	Name 	string
	Origin 	string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	football := Game{
		Name: "Football",
		Origin: "China",
	}

	cricket := Game{
		Name: "Cricket",
		Origin: "England",
	}

	baseball := Game{
		Name: "Baseball",
		Origin: "US",
	}

	gamelist := []Game{ football, cricket, baseball}

	err := tpl.Execute(os.Stdout, gamelist)
	if err != nil {
		log.Fatalln(err)
	}
}