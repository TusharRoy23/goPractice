package main

import (
	"os"
	"log"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// tpl, err := template.ParseGlob("template/*")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	err = tpl.ExecuteTemplate(os.Stdout, "ab.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "cd.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}