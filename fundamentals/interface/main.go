package main

import "fmt"

type person struct {
	Firstname string
	Lastname  string
}

type secretAgent struct {
	person
	active bool
}

func (se secretAgent) kill() {
	fmt.Println("Secret agent's name ", se.Firstname, " & can kill: ", se.active)
}

func (pe person) speak() {
	fmt.Println("Who can speak: ", pe.Lastname)
}

type human interface {
	kill()
}

type ordHuman interface {
	speak()
}

func bar(h human) {
	switch h.(type) { // This is called type assertion in GO
	case secretAgent:
		fmt.Println("Passed value : ", h.(secretAgent).Firstname)
	}
}

func foo(h ordHuman) {
	switch h.(type) { // This is called type assertion in GO
	case person:
		fmt.Println("Passed value : ", h.(person).Lastname)
	}
}

func main() {
	se1 := secretAgent{
		person: person{
			Firstname: "James",
			Lastname:  "Bond",
		},
		active: true,
	}

	pe1 := person{
		Firstname: "Tushar",
		Lastname:  "Roy",
	}

	se1.kill()
	fmt.Println(pe1)

	bar(se1)

	pe1.speak()
	foo(pe1)
}
