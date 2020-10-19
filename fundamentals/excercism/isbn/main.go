package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("Type ISBN number: ")
	var str string
	fmt.Scanf("%s", &str)
	var total int64 = 0
	var formula []int64 = []int64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	var number []int64 = []int64{}
	// formula
	// (x1 * 10 + x2 * 9 + x3 * 8 + x4 * 7 + x5 * 6 + x6 * 5 + x7 * 4 + x8 * 3 + x9 * 2 + x10 * 1) mod 11 == 0
	for _, val := range str {
		match, _ := regexp.MatchString("^[0-9]+$", string(val))
		if match {
			num, _ := strconv.ParseInt(string(val), 10, 64)
			number = append(number, num)
		}
	}

	if len(number) == 10 {
		for index, val := range number {
			total += val * formula[index]
		}

		if total%2 == 0 {
			fmt.Println("Valid ISBN-10")
		} else {
			fmt.Println("Invalid ISBN-10")
		}
	} else {
		fmt.Println("Invalid ISBN-10")
	}
}
