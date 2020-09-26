package main

import "fmt"

func wildCard(binString []string, start int) {

	if len(binString) == start {
		fmt.Println("aa: ", binString)
		return
	}
	if string(binString[start]) == "?" {
		for _, value := range "01" {
			binString[start] = string(value)

			wildCard(binString, start+1)
			fmt.Println("before: ", binString)

			binString[start] = "?"

			fmt.Println("after: ", binString)
		}
		return
	}
	wildCard(binString, start+1)
}

func main() {
	binString := []string{"1", "?", "?"}

	wildCard(binString, 0)

	// Result will be

	// 1001 11 1 1 0
	// 1001 11 1 0 0

	// 1001 00 1 0 0
	// 1001 00 1 1 0

	// 1001 01 1 0 0
	// 1001 01 1 1 0

	// 1001 10 1 0 0
	// 1001 10 1 1 0
}
