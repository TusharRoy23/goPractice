package main

import (
	"fmt"
	"net/http"
	"github.com/satori/go.uuid"
)

func main()  {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request)  {
	cookie, err := req.Cookie("session")
	if err != nil {
		id := uuid.New(4)
		cookie = http.SetCookie(res, &http.Cookie{
			Name: "session",
			Value: id.String(),
			HttpOnly: true,
			Path: "/"
		})
	}
	fmt.Println(cookie)
}