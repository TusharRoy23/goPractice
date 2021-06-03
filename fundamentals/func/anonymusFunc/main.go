package main

import "fmt"

func anon(n int) func() int {
	return func() int {
		n++
		return n
	}
}

func main() {
	value := anon(10)
	fmt.Printf("%T\n", value)
	fmt.Printf("%T\n", value())
	fmt.Printf("%v\n", value())
}
