package main

import "fmt"

func main() {
	x := []int{41, 42, 43, 44, 45, 46, 47, 48, 49, 50}

	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
}
