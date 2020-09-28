package main

import "fmt"

func main() {
	arr := []int{21, 4, 2, 13, 10, 0, 19, 11, 7, 5, 23, 18, 9, 14, 6, 8, 1, 20, 17, 3, 16, 22, 24, 15, 12}
	bubbleSort(arr)
	fmt.Println("final: ", arr)
}

func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		if !swap(arr, i) {
			return
		}
	}
}

func swap(arr []int, sorted int) bool {
	var firstIndex int = 0
	var secondIndex int = 1
	var didSwap bool = false

	for secondIndex < (len(arr) - sorted) {
		firstNumber := arr[firstIndex]
		secondNumber := arr[secondIndex]

		fmt.Println(arr)

		if arr[firstIndex] > arr[secondIndex] {
			arr[firstIndex] = secondNumber
			arr[secondIndex] = firstNumber

			didSwap = true
			fmt.Println("didSwap: ", didSwap)
		}
		firstIndex++
		secondIndex++
	}

	return didSwap
}
