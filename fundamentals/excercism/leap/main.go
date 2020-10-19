package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("Type a year: ")
	var year string
	fmt.Scanf("%s", &year)

	match, _ := regexp.MatchString("^[0-9]+$", string(year))

	if match {
		num, _ := strconv.ParseInt(string(year), 10, 64)
		if num%4 == 0 || num%400 == 0 && num%100 == 0 {
			fmt.Println("This is leap year")
		} else {
			fmt.Println("This is not leap year")
		}
	} else {
		fmt.Println("Only numbers are allowed")
	}
}
