package main

import "fmt"

func main() {
	var numbers []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Difference: ", sqOfSum(numbers)-sumOfSqrs(numbers))
}

func sqOfSum(num []int) int {
	sum := 0
	for _, val := range num {
		sum += val
	}
	return sum * sum
}

func sumOfSqrs(num []int) int {
	sum := 0
	for _, val := range num {
		sum += (val * val)
	}
	return sum
}
