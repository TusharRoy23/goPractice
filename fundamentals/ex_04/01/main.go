package main

import "fmt"

func main() {
	x := [5]int{1, 2, 3, 4, 5}

	for key, value := range x {
		fmt.Println("key-> ", key, "value->", value)
	}

	fmt.Printf("%T\n", x)
}
