package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", setCookie)
	http.HandleFunc("/path-cookie/set", setPathCookie)
	http.HandleFunc("/path-cookie/expire", expirePathCookie)
	http.HandleFunc("/expire", expire)
	http.HandleFunc("/read", read)

	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, `<h1><a href="/set">Set Cookie</a></h1>`)
}

func read(res http.ResponseWriter, req *http.Request) {
	cr, err := req.Cookie("Cookie-Name")
	if err != nil {
		http.Redirect(res, req, "/set", http.StatusSeeOther) // HTTP 303
		return
	}
	fmt.Fprintf(res, `<h1>%s</h1>`, cr)
	fmt.Fprintf(res, `<h1><a href="/expire">Expire</a></h1>`)
}

func setCookie(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "Cookie-Name",
		Value: "ABCD",
		Path:  "/",
	})

	fmt.Fprintln(res, `<h1><a href="/read">Read Cookie</a></h1>`)
}

func setPathCookie(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "Path-Cookie",
		Value: "EFGH",
		Path:  "/path-cookie/set",
	})
	fmt.Fprintln(res, `<h1><a href="/read">Read Path Cookie</a></h1>`)
}

func getPathCookie(res http.ResponseWriter, req *http.Request) {
	cr, err := req.Cookie("Path-Cookie")
	if err != nil {
		http.Redirect(res, req, "/path-cookie/set", http.StatusSeeOther)
		return
	}
	fmt.Fprintf(res, `<h1>%s</h1>`, cr)
	fmt.Fprintf(res, `<h1><a href="/expire">Expire</a></h1>`)
}

func expire(res http.ResponseWriter, req *http.Request) {
	cook, err := req.Cookie("Cookie-Name")
	if err != nil {
		http.Redirect(res, req, "/set", http.StatusSeeOther)
		return
	}

	cook.MaxAge = -1
	http.SetCookie(res, cook)
	http.Redirect(res, req, "/", http.StatusSeeOther)
}

func expirePathCookie(res http.ResponseWriter, req *http.Request) {
	cok, err := req.Cookie("Path-Cookie")
	if err != nil {
		http.Redirect(res, req, "/path-cookie/set", http.StatusSeeOther)
		return
	}
	cok.MaxAge = -1
	http.SetCookie(res, cok)
	http.Redirect(res, req, "/", http.StatusSeeOther)
}
