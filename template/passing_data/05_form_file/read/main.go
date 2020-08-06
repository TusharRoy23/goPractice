package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

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
			return
		}
		defer file.Close()

		// for information
		fmt.Println("\nfile:", file, "\nheader:", header, "\nerr:", err)

		bs, err := ioutil.ReadAll(file) // read file's content
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `
		<form method="POST" enctype="multipart/form-data">
			<input type="file" name="q">
			<input type="submit">
		</form>
		<br>`+s)
}
