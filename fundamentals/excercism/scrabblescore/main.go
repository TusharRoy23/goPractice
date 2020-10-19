package main

import (
	"fmt"
	"strings"
)

func main() {
	var score int
	var scrab string
	fmt.Println("Type a word or letter or sentence: ")
	_, err := fmt.Scanf("%s\n", &scrab)
	if err != nil {
		fmt.Println(err)
	}

	for _, value := range scrab {
		score += getScore(string(value))
	}

	fmt.Println("Score: ", score)
}

func getScore(str string) int {
	str = strings.ToUpper(str)
	if str == "A" || str == "E" || str == "I" || str == "O" || str == "U" || str == "L" || str == "N" || str == "R" || str == "S" || str == "T" {
		return 1
	} else if str == "D" || str == "G" {
		return 2
	} else if str == "B" || str == "C" || str == "M" || str == "P" {
		return 3
	} else if str == "F" || str == "H" || str == "V" || str == "W" || str == "Y" {
		return 4
	} else if str == "K" {
		return 5
	} else if str == "J" || str == "X" {
		return 8
	} else if str == "Q" || str == "Z" {
		return 10
	} else {
		return 0
	}
}
