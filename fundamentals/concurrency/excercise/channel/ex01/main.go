package main

import "fmt"

func main() {
	// Unbuffered channel
	var c1 chan float64

	// Unbuffered one way channels
	c2 := make(chan<- rune)
	c3 := make(<-chan rune)

	// Bidirectional buffered channel
	c4 := make(chan float64, 10)

	fmt.Printf("%T %T %T %T\n", c1, c2, c3, c4)

	c5 := make(chan string)
	go func(n string) {
		c5 <- n
	}("Golang")

	value := <-c5
	fmt.Println("Value: ", value)
}
