package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", 48)
	if err != nil {
		log.Fatalln("err")
	}

	// newFile, err := os.Create("index.html")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// defer newFile.Close()

	// err = tpl.Execute(newFile, nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
}