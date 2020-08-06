package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("views/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	var s string
	if req.Method == http.MethodPost {
		// Open File
		file, header, err := req.FormFile("q")
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
		defer file.Close()

		// for information
		fmt.Println("\nfile:", file, "\nheader:", header, "\nerr", err)

		// read file's content
		// bs = byte string
		bs, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)

		// store on server
		dest, err := os.Create(filepath.Join("./user/", header.Filename))
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dest.Close()

		// vt, err := dest.Write(bs) // bs = byte string
		// fmt.Println("\n check: ", vt)
		// if err != nil {
		// 	http.Error(res, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		// Write into another file
		err = ioutil.WriteFile("./user/test.txt", bs, 0644)
		if err != nil {
			log.Fatalln(err)
		}
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl.ExecuteTemplate(res, "index.gohtml", s)
}
