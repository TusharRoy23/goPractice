package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.jpg", dogImg)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Foo ran")
}

func dog(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(res, "index.gohtml", nil)
}

func dogImg(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "toby.jpg")
}
