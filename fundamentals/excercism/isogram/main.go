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
		var count int
		sentence := scanner.Text()
		for index, value := range sentence {
			for ind, inVal := range sentence {
				if index != ind {
					if string(value) == string(inVal) && string(inVal) != " " && string(inVal) != "-" {
						count++
						break
					}
				}
			}
		}

		if count == 0 {
			fmt.Println("This is Isogram")
		} else {
			fmt.Println("This is not Isogram")
		}
	}
}
