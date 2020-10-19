package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("Type the numbers: ")
	var str string
	fmt.Scanf("%s", &str)
	fmt.Println("Series of digit: ")
	var number int
	fmt.Scanf("%d", &number)

	match, _ := regexp.MatchString("^[0-9]+$", str)

	if match {
		if len(str) <= number {
			fmt.Println(str)
		} else {
			var i int = number
			for index := range str {

				if len(str[index:]) >= number {
					fmt.Println("array ", index, ": ", str[index:i])
				} else {
					break
				}
				i++
			}
		}
	} else {
		fmt.Println("Only numbers")
	}
}
