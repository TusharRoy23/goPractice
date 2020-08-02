package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type item struct {
	Name  string
	Price int
}

type meal struct {
	Meal string
	Item []item
}

type menu []meal

type restaurant struct {
	Name string
	Menu menu
}

type restaurants []restaurant

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	value := restaurants{
		restaurant{
			Name: "Restaurant One",
			Menu: menu{
				meal{
					Meal: "Breakfast",
					Item: []item{
						item{
							Name:  "Ruti",
							Price: 10,
						},
						item{
							Name:  "Bhaji",
							Price: 15,
						},
					},
				},
				meal{
					Meal: "Lunch",
					Item: []item{
						item{
							Name:  "Rice",
							Price: 10,
						},
						item{
							Name:  "Daal",
							Price: 15,
						},
					},
				},
			},
		},
		restaurant{
			Name: "Restaurant Two",
			Menu: menu{
				meal{
					Meal: "Breakfast",
					Item: []item{
						item{
							Name:  "Khichuri",
							Price: 50,
						},
						item{
							Name:  "Porota",
							Price: 15,
						},
					},
				},
				meal{
					Meal: "Lunch",
					Item: []item{
						item{
							Name:  "Biriyani",
							Price: 100,
						},
						item{
							Name:  "Chicken",
							Price: 80,
						},
					},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, value)
	if err != nil {
		log.Fatalln(err)
	}
}
