package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func defaultRouter(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "This is default router")
}

func dogRouter(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "This is Dog Router")
}

func myselfRouter(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = tpl.ExecuteTemplate(res, "index.gohtml", "Tushar")
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func main() {

	http.Handle("/", http.HandlerFunc(defaultRouter))
	http.HandleFunc("/dog/", http.HandlerFunc(dogRouter))
	http.HandleFunc("/me/", http.HandlerFunc(myselfRouter))

	http.ListenAndServe(":8080", nil)
}
