package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 6, 10)
	fmt.Println(cap(slice))

	slice = append(slice, 1, 2, 3, 4, 7, 8, 9, 10, 11, 11, 11, 11, 11, 11, 12, 13, 14)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	s := slice[2:]
	s = append(s, 9, 3)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}
