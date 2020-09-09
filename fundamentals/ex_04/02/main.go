package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for key, value := range x {
		fmt.Println("key-> ", key, "value->", value)
	}

	fmt.Printf("%T\n", x)
}
