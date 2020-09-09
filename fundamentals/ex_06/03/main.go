package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

func (p person) speak() {
	fmt.Println("Name:", p.First, "& Age:", p.Age)
}

func main() {
	p1 := person{
		"Tushar", "Roy", 28,
	}

	p1.speak()
}
