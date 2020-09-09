package main

import "fmt"

func main() {

	g := func(n ...int) int {
		total := 0
		for _, value := range n {
			total += value
		}
		return total
	}

	st := foo(g, []int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(st)
}

func foo(f func(xi ...int) int, in []int) int {
	return f(in...)
}
