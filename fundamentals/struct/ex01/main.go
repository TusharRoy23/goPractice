package main

import (
	"fmt"
)

func main() {
	type person struct {
		name   string
		age    int
		colors []string
	}

	me := person{
		name:   "Tushar",
		age:    28,
		colors: []string{"red", "green", "blue"},
	}

	you := person{
		name:   "John",
		age:    39,
		colors: []string{"purple", "pink", "yellow"},
	}

	me.name = "Andrei"

	var colors []string = you.colors
	colors = append(colors, "black")
	you.colors = colors

	// iterating colors

	for _, value := range you.colors {
		fmt.Println("Colors: ", value)
	}

	fmt.Printf("%v\n", me)
	fmt.Printf("%v\n", you)
}
