package main

import "fmt"

type person struct {
	Firstname string
	Lastname  string
}

func main() {
	p1 := person{
		"Tushar", "Roy",
	}
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(p *person) {
	(*p).Firstname = "Janeva"
	(*p).Lastname = "Yor"
	p.Firstname = "Rahsut" // short hand for line 19
}
