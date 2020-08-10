package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/multi", multiCookie)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(res http.ResponseWriter, req *http.Request) {
	// http.SetCookie(response, *cookie Cookie)
	http.SetCookie(res, &http.Cookie{
		Name:  "my-cookie",
		Value: "Some valuez",
		Path:  "/",
	})
	fmt.Fprintln(res, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(res, "in chrome go to: dev tools / application / cookies")
}

func read(res http.ResponseWriter, req *http.Request) {
	cook, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(res, http.StatusText(400), http.StatusBadRequest)
		return
	}
	fmt.Fprintln(res, "Your Cookie:", cook)

	c2, err := req.Cookie("cookie-2")
	if err != nil {
		http.Error(res, http.StatusText(400), http.StatusBadRequest)
		return
	}
	fmt.Fprintln(res, "Your Cookie:", c2)

	c3, err := req.Cookie("cookie-3")
	if err != nil {
		http.Error(res, http.StatusText(400), http.StatusBadRequest)
		return
	}
	fmt.Fprintln(res, "Your Cookie:", c3)

	c4, err := req.Cookie("count-visit")
	if err != nil {
		http.Error(res, http.StatusText(400), http.StatusBadRequest)
		return
	}
	fmt.Fprintln(res, "Your Cookie:", c4)
}

func multiCookie(res http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("count-visit")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "count-visit",
			Value: "0",
		}
	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(res, cookie)

	http.SetCookie(res, &http.Cookie{
		Name:  "cookie-2",
		Value: "Value 2",
	})
	http.SetCookie(res, &http.Cookie{
		Name:  "cookie-3",
		Value: "Value 3",
	})

	fmt.Fprintln(res, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(res, "in chrome go to: dev tools / application / cookies")
}
