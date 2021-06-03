package main

import "fmt"

func vardic(sum ...int) int {
	total := 0
	for _, value := range sum {
		total += value
	}
	return total
}

func main() {
	val := vardic(1, 2, 3, 4, 5, 6)
	fmt.Println(val)
}
