package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("views/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about/", about)
	http.HandleFunc("/contact/", contact)
	http.HandleFunc("/apply/", apply)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "index.gohtml", nil)
	HandleError(res, err)
}

func about(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "about.gohtml", nil)
	HandleError(res, err)
}

func contact(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "contact.gohtml", nil)
	HandleError(res, err)
}

func apply(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		err := tpl.ExecuteTemplate(res, "applyProcess.gohtml", nil)
		HandleError(res, err)
		return
	}

	err := tpl.ExecuteTemplate(res, "apply.gohtml", nil)
	HandleError(res, err)
}

func HandleError(res http.ResponseWriter, err error) {
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
