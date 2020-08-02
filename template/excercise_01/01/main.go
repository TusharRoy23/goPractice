package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type hotel struct {
	Name, Address, City, Zip, Region string
}

type region struct {
	Region string
	Hotels []hotel
}

type Regions []region

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	value := Regions{
		region{
			Region: "Southern",
			Hotels: []hotel{
				hotel{
					Name:    "Hotel California",
					Address: "42 Sunset Boulevard",
					City:    "Los Angeles",
					Zip:     "95612",
				},
				hotel{
					Name:    "H",
					Address: "4",
					City:    "L",
					Zip:     "95612",
				},
			},
		},
		region{
			Region: "Northen",
			Hotels: []hotel{
				hotel{
					Name:    "Hotel California",
					Address: "42 Sunset Boulevard",
					City:    "Los Angeles",
					Zip:     "95612",
				},
				hotel{
					Name:    "H",
					Address: "4",
					City:    "L",
					Zip:     "95612",
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, value)
	if err != nil {
		log.Fatalln(err)
	}
}
