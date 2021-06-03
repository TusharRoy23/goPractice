package main

import "fmt"

func nakedFunc(n ...int) (v int) {
	for _, value := range n {
		v += value
	}
	return
}

func main() {
	v := nakedFunc(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(v)
}
