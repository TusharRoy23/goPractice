package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

type session struct {
	un           string
	lastActivity time.Time
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]session{} // map[keyType]valueType
var dbSessionsCleaned time.Time

const sessionLength int = 30

func init() {
	tpl = template.Must(template.ParseGlob("views/*"))
	dbSessionsCleaned = time.Now()
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	showSessions() // for demonstration purposes
	tpl.ExecuteTemplate(res, "index.gohtml", u)
}

func bar(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	if !alreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	if u.Role != "007" {
		http.Error(res, "Need 007 role as a user", http.StatusForbidden)
		return
	}
	showSessions() // for demonstration purposes
	tpl.ExecuteTemplate(res, "bar.gohtml", u)
}

func signup(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	var u user
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")

		if _, ok := dbUsers[un]; ok {
			http.Error(res, "user already exists !", http.StatusForbidden)
			return
		}
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "sesion",
			Value: sID.String(),
		}
		c.MaxAge = sessionLength
		http.SetCookie(res, c)
		dbSessions[c.Value] = session{un, time.Now()}

		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(res, "Internal server error", http.StatusInternalServerError)
			return
		}
		u = user{un, bs, f, l, r}
		dbUsers[un] = u
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	showSessions() // for demonstration purposes
	tpl.ExecuteTemplate(res, "signup.gohtml", u)
}

func login(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		u, ok := dbUsers[un]
		if !ok {
			http.Error(res, "Username and/or Password do not match", http.StatusForbidden)
			return
		}

		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(res, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(res, c)
		dbSessions[c.Value] = session{un, time.Now()}
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	showSessions() // for demonstration purposes

	tpl.ExecuteTemplate(res, "login.gohtml", nil)
}

func logout(res http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
	}
	c, _ := req.Cookie("session")
	delete(dbSessions, c.Value) // delete(m map[Type]Type1, key Type)
	fmt.Println("print:", dbSessions)
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(res, c)
	// clean up dbSessions
	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) {
		go cleanSessions()
	}
	http.Redirect(res, req, "/login", http.StatusSeeOther)
}
