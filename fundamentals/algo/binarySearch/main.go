package main

import "fmt"

func main() {
	var arr []int = []int{2, 5, 6, 3, 10, 1}

	index := binarySearch(arr, 2)
	if index >= 0 {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
	}
}

func binarySearch(list []int, search int) int {
	var low int = 0
	var high int = len(list) - 1

	for low <= high {
		var mid int = low + (high-low)/2
		var midValue = list[mid]

		if midValue == search {
			return mid
		} else if midValue > search {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
