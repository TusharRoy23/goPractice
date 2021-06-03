package main

import "fmt"

func f1(a uint) (int, int) {
	fact := 1
	sum := 0
	for i := 1; i <= int(a); i++ {
		fact *= i
		sum += i
	}
	return fact, sum
}

func main() {
	fact, sum := f1(4)

	fmt.Printf("fact %v\n", fact)
	fmt.Printf("sum %v\n", sum)
}
