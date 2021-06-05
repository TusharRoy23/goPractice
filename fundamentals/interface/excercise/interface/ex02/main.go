package main

import "fmt"

func main() {
	var empty interface{}
	fmt.Printf("%T\n", empty)

	empty = 4
	fmt.Printf("%T\n", empty)

	empty = 4.4
	fmt.Printf("%T\n", empty)

	empty = []int{1, 2, 3}
	fmt.Printf("%T\n", empty)

	empty = append(empty.([]int), 10)
	fmt.Printf("%v\n", empty)
}
