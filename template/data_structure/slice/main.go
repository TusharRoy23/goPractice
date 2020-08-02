package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	alphabets := []string{"AB", "CD", "EF", "GH"}

	err := tpl.Execute(os.Stdout, alphabets)
	if err != nil {
		log.Fatalln(err)
	}
}