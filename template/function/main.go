package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type Country struct {
	Name string
	Capital string
}

type Language struct {
	Origin string
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	// tpl = template.Must(template.ParseFiles("tpl.gohtml"))
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func main() {
	Bangladesh := Country {
		Name: "Bangladesh",
		Capital: "Dhaka",
	}

	India := Country {
		Name: "India",
		Capital: "Delhi",
	}

	Nepal := Country {
		Name: "Nepal",
		Capital: "Katmundu",
	}

	Bangla := Language {
		Origin: "Indian Sub-continent",
	}

	Hindi := Language {
		Origin: "India",
	}

	English := Language {
		Origin: "England",
	}

	countries := []Country{ Bangladesh, India, Nepal }
	languages := []Language{ Bangla, Hindi, English }

	data := struct {
		Desh 	[]Country
		Vasha 	[]Language
	}{
		countries,
		languages,
	}
	
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}