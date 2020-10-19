package main

import "fmt"

func main() {
	var numbers []int = []int{3, 5, 7}

	fmt.Println("Type a number: ")
	var inp int
	_, err := fmt.Scanf("%d\n", &inp)

	if err != nil {
		fmt.Println(err)
	}

	var str string = ""

	for i := 0; i < len(numbers); i++ {
		if int(inp%numbers[i]) == 0 && int(numbers[i]) == 3 {
			str += "Pling"
		} else if int(inp%numbers[i]) == 0 && int(numbers[i]) == 5 {
			str += "Plang"
		} else if int(inp%numbers[i]) == 0 && int(numbers[i]) == 7 {
			str += "Plong"
		}
	}

	if str != "" {
		fmt.Println(str)
	} else {
		fmt.Println(inp)
	}
}
