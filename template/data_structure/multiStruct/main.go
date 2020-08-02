package main

import (
	"os"
	"log"
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

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
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
	
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}