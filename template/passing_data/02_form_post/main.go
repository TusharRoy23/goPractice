package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	value := req.FormValue("q")
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `
		<form method="POST">
		 <input type="text" name="q">
		 <input type="submit">
		</form>
	`+value)
}
