package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any hotdog type functions will be server Handler")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
