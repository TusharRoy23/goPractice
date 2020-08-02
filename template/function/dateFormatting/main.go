package main

import (
	"os"
	"text/template"
	"log"
	"time"
	"fmt"
)

var tpl *template.Template

var fm = template.FuncMap{
	"fdate": monthDayYear,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}