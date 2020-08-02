package main

import (
	"os"
	"log"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	alphabates := map[string]string{
		"1st":"AB",
		"2nd":"CD",
		"3rd":"EF",
	}
	err := tpl.Execute(os.Stdout, alphabates)
	if err != nil {
		log.Fatalln(err)
	}
}