package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/authenticate", auth)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		c = &http.Cookie{
			Name:  "session",
			Value: "",
		}
	}

	if req.Method == http.MethodPost {
		em := req.FormValue("email")
		c.Value = em + `|` + getCode(em+"s")
	}

	http.SetCookie(res, c)

	io.WriteString(res, `<!DOCTYPE html>
	<html>
	  <body>
	    <form method="POST">
	      <input type="email" name="email">
	      <input type="submit">
	    </form>
	    <a href="/authenticate">Validate This `+c.Value+`</a>
	  </body>
	</html>`)
}

func auth(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(res, req, "/", http.StatusSeeOther)
	}

	if c.Value == "" {
		http.Redirect(res, req, "/", http.StatusSeeOther)
	}

	xs := strings.Split(c.Value, "|")
	email := xs[0]
	codeReceive := xs[1]
	codeCheck := getCode(email)

	if codeReceive != codeCheck {
		fmt.Println("HMAC codes didn't match !")
		fmt.Println("codeReceive: ", codeReceive)
		fmt.Println("codeCheck: ", codeCheck)
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	io.WriteString(res, `<!DOCTYPE html>
	<html>
	  <body>
	  	<h1>`+codeReceive+` - RECEIVED </h1>
	  	<h1>`+codeCheck+` - RECALCULATED </h1>
	  </body>
	</html>`)
}

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("secretKey"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}
