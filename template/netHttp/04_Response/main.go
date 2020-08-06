package main

import (
	"fmt"
	"net/http"
)

type hotdog int

// This function implements http Handler Interface
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Mcleod-Key", "This is from mcload")
	w.Header().Set("Secret-key", "ABCD")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>This is Header !</h1>")
}

func main() {
	var handleFunc hotdog
	http.ListenAndServe(":8080", handleFunc)
}
