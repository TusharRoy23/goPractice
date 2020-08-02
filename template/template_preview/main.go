package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("views/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public", fs))
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "index.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
