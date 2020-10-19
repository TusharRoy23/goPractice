package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Type a sentence: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		sentence := scanner.Text()

		length := len(sentence) - 1

		// using for-loop
		// for i := length; i >= 0; i-- {
		// 	fmt.Print(string(sentence[i]))
		// }
		// fmt.Println()

		// recursive function
		fmt.Println(reverse(sentence, length))
	}
}

func reverse(str string, length int) string {
	if length == 0 {
		return string(str[0])
	}
	return string(str[length]) + reverse(str, length-1)
}
