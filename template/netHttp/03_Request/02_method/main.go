package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hotdog int

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

// This function implements http Handler Interface
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method      string
		Submissions url.Values
	}{
		r.Method,
		r.Form,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
