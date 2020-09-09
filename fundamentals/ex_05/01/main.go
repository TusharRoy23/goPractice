package main

import "fmt"

type person struct {
	Firstname string
	Lastname  string
	Icecream  []string
}

func main() {
	p1 := person{
		Firstname: "Tushar",
		Lastname:  "Roy",
		Icecream: []string{
			"fav1",
			"fav2",
			"fav3",
		},
	}

	p2 := person{
		Firstname: "Risen",
		Lastname:  "Boy",
		Icecream: []string{
			"lav1",
			"lav2",
			"lav3",
		},
	}

	fmt.Println(p1.Firstname, "-", p2.Firstname)

	for _, value := range p1.Icecream {
		fmt.Println(value)
	}

	for _, value := range p2.Icecream {
		fmt.Println(value)
	}
}
