package main

import "fmt"

type person struct {
	Firstname string
	Lastname  string
}

func main() {
	arr := []person{
		person{"Jon", "Calhoun"}, person{"Bob", "Smith"}, person{"John", "Smith"},
		person{"Larry", "Green"}, person{"Joseph", "Page"}, person{"George", "Costanza"},
		person{"Jerry", "Seinfeld"}, person{"Shane", "Calhoun"},
	}

	bubbleSort(arr)

	for _, value := range arr {
		fmt.Println(value.Firstname, value.Lastname)
	}
}

func bubbleSort(arr []person) {
	for i := 0; i < len(arr); i++ {
		if !swap(arr, i) {
			return
		}
	}
}

func swap(arr []person, sorted int) bool {
	var firstIndex int = 0
	var secondIndex int = 1
	var didSwap bool = false

	for (len(arr) - sorted) > secondIndex {
		firstName := arr[firstIndex]
		secondName := arr[secondIndex]

		if firstName.Lastname == secondName.Lastname {
			if firstName.Firstname > secondName.Firstname {
				arr[firstIndex] = secondName
				arr[secondIndex] = firstName

				didSwap = true
			}
		} else {
			if firstName.Lastname > secondName.Lastname {
				arr[firstIndex] = secondName
				arr[secondIndex] = firstName

				didSwap = true
			}
		}

		firstIndex++
		secondIndex++
	}
	return didSwap
}
