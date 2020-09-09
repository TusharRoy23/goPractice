package main

import "fmt"

func foo(n ...int) int {
	total := 0
	for _, value := range n {
		total += value
	}
	return total
}

func bar(n []int) int {
	total := 0
	for _, value := range n {
		total += value
	}
	return total
}

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8}
	val := foo(x...)
	fmt.Println(val)
	val2 := bar(x)
	fmt.Println(val2)
}
