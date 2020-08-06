package main

import (
	"io"
	"net/http"
)

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Dog dogs")
}

func cat(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Cat cat")
}

func main() {
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/cat", http.HandlerFunc(cat))

	http.ListenAndServe(":8080", nil)
}
