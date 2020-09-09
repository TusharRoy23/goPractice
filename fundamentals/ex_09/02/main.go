package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func (p *person) speak() string {
	fmt.Println("who am I ?")
	return p.Name
}

type human interface {
	speak() string
}

func saySomething(h human) {
	fmt.Println("What is your name ? ", h.speak())
}

func main() {
	p1 := person{"Tushar Roy", 28}
	saySomething(&p1)
	p1.speak()
}
