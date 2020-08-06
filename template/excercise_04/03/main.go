package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	fs := http.StripPrefix("/resources", http.FileServer(http.Dir("./public")))
	http.Handle("/resources/", fs)
	http.HandleFunc("/", dogs)
	http.ListenAndServe(":8080", nil)
}

func dogs(res http.ResponseWriter, req *http.Request) {
	// err := tpl.Execute(res, nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	tpl.ExecuteTemplate(res, "index.gohtml", nil)
}
