package main

import "fmt"

func main() {
	type grades struct {
		grade  int
		course string
	}

	type person struct {
		name  string
		grade grades
	}

	me := person{
		name: "Tushar",
		grade: grades{
			grade:  100,
			course: "Math",
		},
	}

	you := person{
		name: "John",
		grade: grades{
			grade:  100,
			course: "English",
		},
	}

	you.grade.grade = 98
	you.grade.course = "Golang"

	fmt.Printf("%v\n", me)
	fmt.Printf("%v\n", you)
}
