package main

import "fmt"

func main() {
	var A []int = []int{1, 2, 3, 4, 5, 6}
	var B []int = []int{2, 3, 4}
	var motherArr []int
	var childArr []int

	if len(A) >= len(B) {
		motherArr = A
		childArr = B
	} else {
		motherArr = B
		childArr = A
	}

	for index, value1 := range motherArr {
		if value1 == childArr[0] {
			motherArr = motherArr[index:]
			break
		}
	}

	var start int = 0
	var found int = 0
	for _, val1 := range motherArr {
		for i := start; i < len(childArr); i++ {
			if val1 == childArr[i] {
				start = i + 1
				i = len(childArr)
				found++
			}
		}
	}

	if found == len(childArr) {
		if len(A) == len(B) {
			fmt.Println("A is equal to B")
		} else if len(A) > len(B) {
			fmt.Println("B is a sublist of A")
		} else if len(A) < len(B) {
			fmt.Println("A is a sublist of B")
		} else {
			fmt.Println("Found")
		}

	} else {
		fmt.Println("A is not a superlist of, sublist of or equal to B", found)
	}
}
