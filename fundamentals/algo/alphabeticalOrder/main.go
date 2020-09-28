package main

import "fmt"

func main() {
	arr := []string{"dog", "cat", "alligator", "cheetah", "rat", "moose", "cow", "mink", "porcupine", "dung beetle", "camel", "steer", "bat", "hamster", "horse", "colt", "bald eagle", "frog", "rooster"}
	alphaSort(arr)
	fmt.Println(arr)
}

func alphaSort(arr []string) {
	for i := 0; i < len(arr); i++ {
		if !swap(arr, i) {
			return
		}
	}
}

func swap(arr []string, sorted int) bool {
	var firstIndex int = 0
	var secondIndex int = 1
	var didSwap bool = false

	for (len(arr) - sorted) > secondIndex {
		firstNumber := arr[firstIndex]
		secondNumber := arr[secondIndex]

		if firstNumber > secondNumber {
			arr[firstIndex] = secondNumber
			arr[secondIndex] = firstNumber

			didSwap = true
		}
		firstIndex++
		secondIndex++
	}
	return didSwap
}
