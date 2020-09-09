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

	m := map[string]person{
		p1.Lastname: p1,
		p2.Lastname: p2,
	}

	for key, value := range m {
		fmt.Println(value)
		fmt.Println("---------")
		fmt.Println(key)
	}
}
