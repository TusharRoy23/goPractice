package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(res, `<img src="./toby.jpg">`)
	io.WriteString(res, "hell ya")
}

func dogPic(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "tody.jpg")
}
