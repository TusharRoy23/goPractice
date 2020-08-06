package main

import (
	"io"
	"net/http"
)

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Dog dog")
}

func cat(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Cat cat")
}

func main() {
	// HandleFunc is equivalent to ServeHTTP to function
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/cat", cat)

	http.ListenAndServe(":8080", nil)
}
