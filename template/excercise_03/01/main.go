package main

import (
	"io"
	"net/http"
)

func defaultRouter(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "This is default router")
}

func dogRouter(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "This is Dog Router")
}

func myselfRouter(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello Tushar Roy ")
}

func main() {

	http.HandleFunc("/", defaultRouter)
	http.HandleFunc("/dog/", dogRouter)
	http.HandleFunc("/me/", myselfRouter)

	http.ListenAndServe(":8080", nil)
}
